// SPDX-License-Identifier: MIT-open-group
pragma solidity ^0.8.0;
import "@openzeppelin/contracts/token/ERC20/ERC20.sol";
import "@openzeppelin/contracts-upgradeable/proxy/utils/Initializable.sol";
import "@openzeppelin/contracts-upgradeable/token/ERC20/ERC20Upgradeable.sol";
import "contracts/MadByte.sol";
import "contracts/EventEmitter.sol";
import "hardhat/console.sol";

contract BridgePool is Initializable {
    mapping(address => uint256) internal _depositMB;
    mapping(address => uint256) internal _sidechainBurnedMB;
    address internal _eventEmitter;
    address internal _madByte;

    function initialize(address madByte_, address eventEmitter_) public initializer {
        _eventEmitter = eventEmitter_;
        _madByte = madByte_;
    }

    function deposit(uint256 amountMB_) public returns (uint256) {
        return _deposit(msg.sender, amountMB_);
    }

    function requestWithdrawal(uint256 depositID_) public {
        _requestWithdrawal(msg.sender, depositID_);
    }

    function confirmProofOfBurn(address account_, uint256 amountMB_) public {
        //TODO: set access control.
        _sidechainBurnedMB[account_] += amountMB_;
    }

    function withdraw(uint256 amountMB_) public {
        _withdraw(msg.sender, amountMB_);
    }

    function _deposit(address account_, uint256 amountMB_) internal returns (uint256) {
        ERC20(_madByte).transferFrom(account_, address(this), amountMB_);
        uint256 amountETH = MadByte(_madByte).burn(amountMB_, 0);
        require(amountETH > 0, "BridgePool: Could not burn tokens for deposit");
        _depositMB[account_] += amountETH;
        _emitDepositEvent(account_, amountMB_);
        return 0;
    }

    function _requestWithdrawal(address account_, uint256 amountMB_) internal {
        _emitWithdrawalRequestEvent(account_, amountMB_);
    }

    function _withdraw(address account_, uint256 amountMB_) internal {
        require(
            amountMB_ <= _sidechainBurnedMB[account_],
            "BridgePool: Withdrawal amountMB greater than available balance"
        );
        _sidechainBurnedMB[account_] -= amountMB_;
        // uint256 amountETH = _depositMB[account_];
        uint256 poolBalance_ = MadByte(_madByte).getPoolBalance();
        uint256 totalSupply_ = ERC20(_madByte).totalSupply();
        uint256 amountETH = MadByte(_madByte).madByteToEth(poolBalance_, totalSupply_, amountMB_);
        uint256 madBytes = MadByte(_madByte).mintTo{value: amountETH}(account_, 0);
        console.log(amountMB_, amountETH, madBytes);
        _emitDistributeEvent(account_, madBytes);
        delete _depositMB[account_];
    }

    function _emitDepositEvent(address account_, uint256 amountMB_) internal {
        bytes memory encodedEvent = abi.encode("BridgePool", "deposit", account_, amountMB_);
        EventEmitter(_eventEmitter).emitGenericEvent(encodedEvent);
    }

    function _emitWithdrawalRequestEvent(address account_, uint256 amountMB_) internal {
        bytes memory encodedEvent = abi.encode(
            "BridgePool",
            "requestWithdrawal",
            account_,
            amountMB_
        );
        EventEmitter(_eventEmitter).emitGenericEvent(encodedEvent);
    }

    function _emitDistributeEvent(address account_, uint256 amountMB_) internal {
        bytes memory encodedEvent = abi.encode("BridgePool", "withdrawal", account_, amountMB_);
        EventEmitter(_eventEmitter).emitGenericEvent(encodedEvent);
    }

    receive() external payable {}
}
