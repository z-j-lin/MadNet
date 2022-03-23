import { SignerWithAddress } from "@nomiclabs/hardhat-ethers/signers";
import { BigNumberish } from "ethers";
import { ethers } from "hardhat";
import { ValidatorPoolMock } from "../../typechain-types";
import { expect } from "../chai-setup";
import { Fixture, getFixture } from "../setup";

describe("ValidatorNFT: Tests ValidatorNFT Business Logic methods", async () => {
  let fixture: Fixture;
  let notAdminSigner: SignerWithAddress;
  let adminSigner: SignerWithAddress;
  let amount: BigNumberish;
  let validatorPool: ValidatorPoolMock;

  beforeEach(async function () {
    fixture = await getFixture(true, true);
    const [admin, notAdmin] = fixture.namedSigners;
    adminSigner = await ethers.getSigner(admin.address);
    notAdminSigner = await ethers.getSigner(notAdmin.address);
    validatorPool = fixture.validatorPool as ValidatorPoolMock;
    amount = await validatorPool.getStakeAmount();
    await fixture.madToken.approve(validatorPool.address, amount);
  });

  it("Should mint a token and sender should be the payer and owner", async function () {
    const madBalanceBefore = await fixture.madToken.balanceOf(
      adminSigner.address
    );
    const nftBalanceBefore = await fixture.validatorNFT.balanceOf(
      validatorPool.address
    );
    const rcpt = await (await validatorPool.mintValidatorNFT()).wait();
    expect(rcpt.status).to.be.equal(1);
    expect(await fixture.validatorNFT.ownerOf(1)).to.be.eq(
      validatorPool.address
    );
    expect(
      await fixture.validatorNFT // NFT +1
        .balanceOf(validatorPool.address)
    ).to.equal(nftBalanceBefore.add(1));
    expect(
      await fixture.madToken // MAD -= amount
        .balanceOf(adminSigner.address)
    ).to.equal(madBalanceBefore.sub(amount));
  });

  it("Should burn a token and sender should receive funds", async function () {
    let tx = await validatorPool.mintValidatorNFT();
    const rcpt = await tx.wait();
    expect(rcpt.status).to.be.equal(1);
    const nftBalanceBefore = await fixture.validatorNFT.balanceOf(
      validatorPool.address
    );
    const madBalanceBefore = await fixture.madToken.balanceOf(
      validatorPool.address
    );
    tx = await validatorPool.burnValidatorNFT(1);
    expect((await tx.wait()).status).to.be.equal(1);
    expect(await fixture.madToken.balanceOf(validatorPool.address)).to.be.eq(
      amount
    );
    expect(
      await fixture.validatorNFT // NFT -1
        .balanceOf(validatorPool.address)
    ).to.equal(nftBalanceBefore.sub(1));
    expect(
      await fixture.madToken // MAD +=amount
        .balanceOf(validatorPool.address)
    ).to.equal(madBalanceBefore.add(amount));
  });

  it("Should mint a token to an address and send staking funds from sender address", async function () {
    const madBalanceBefore = await fixture.madToken.balanceOf(
      adminSigner.address
    );
    const nftBalanceBefore = await fixture.validatorNFT.balanceOf(
      notAdminSigner.address
    );
    const rcpt = await (
      await validatorPool.mintToValidatorNFT(notAdminSigner.address)
    ).wait();
    expect(rcpt.status).to.be.equal(1);
    expect(
      await fixture.validatorNFT // NFT +1
        .balanceOf(notAdminSigner.address)
    ).to.equal(nftBalanceBefore.add(1));
    expect(
      await fixture.madToken // MAD -= amount
        .balanceOf(adminSigner.address)
    ).to.equal(madBalanceBefore.sub(amount));
  });

  it("Should burn a token from an address and return staking funds", async function () {
    let tx = await validatorPool.mintValidatorNFT();
    const rcpt = await tx.wait();
    expect(rcpt.status).to.be.equal(1);
    const nftBalanceBefore = await fixture.validatorNFT.balanceOf(
      validatorPool.address
    );
    const madBalanceBefore = await fixture.madToken.balanceOf(
      notAdminSigner.address
    );
    tx = await validatorPool.burnToValidatorNFT(1, notAdminSigner.address);
    expect((await tx.wait()).status).to.be.equal(1);
    expect(
      await fixture.validatorNFT // NFT -1
        .balanceOf(validatorPool.address)
    ).to.equal(nftBalanceBefore.sub(1));
    expect(
      await fixture.madToken // MAD +=amount
        .balanceOf(notAdminSigner.address)
    ).to.equal(madBalanceBefore.add(amount));
  });

  it("Should return correct token uri", async function () {
    let tx = await validatorPool.mintValidatorNFT();
    await tx.wait();

    const tokenId = 1;
    const positionData = await fixture.validatorNFT.getPosition(tokenId);

    const svg =
      `<svg width="500" height="500" viewBox="0 0 500 500" xmlns="http://www.w3.org/2000/svg" xmlns:xlink='http://www.w3.org/1999/xlink'>` +
      `<text x='10' y='20'>Shares: ${positionData.shares.toString()}</text>` +
      `<text x='10' y='40'>Free after: ${positionData.freeAfter.toString()}</text>` +
      `<text x='10' y='60'>Withdraw Free After: ${positionData.withdrawFreeAfter.toString()}</text>` +
      `<text x='10' y='80'>Accumulator (ETH): ${positionData.accumulatorEth.toString()}</text>` +
      `<text x='10' y='100'>Accumulator (Token): ${positionData.accumulatorToken.toString()}</text></svg>`;

    const tokenUriJson =
      `{"name":"MadNET Staked token for position #1", ` +
      `"description":"This NFT represents a staked position on MadNET.` +
      `\\nThe owner of this NFT can modify or redeem the position.` +
      `\\n Shares: ${positionData.shares.toString()}` +
      `\\nFree After: ${positionData.freeAfter.toString()}` +
      `\\nWithdraw Free After: ${positionData.withdrawFreeAfter.toString()}` +
      `\\nAccumulator Eth: ${positionData.accumulatorEth.toString()}` +
      `\\nAccumulator Token: ${positionData.accumulatorToken.toString()}` +
      `\\nToken ID: ${tokenId.toString()}", ` +
      `"image": "data:image/svg+xml;base64,${btoa(svg)}"}`;

    const expectedTokenUriData = `data:application/json;base64,${btoa(
      tokenUriJson
    )}`;

    const tokenUri = await fixture.validatorNFT.tokenURI(tokenId);

    const parsedJson = JSON.parse(
      atob(tokenUri.replace("data:application/json;base64,", ""))
    );

    await expect(tokenUri).to.be.equal(expectedTokenUriData);
    await expect(
      atob(parsedJson.image.replace("data:image/svg+xml;base64,", ""))
    ).to.be.equal(svg);
  });
});
