// SPDX-License-Identifier: MIT-open-group
pragma solidity ^0.8.11;

abstract contract StakingNFTStorage {
    // Position describes a staked position
    //
    // TODO: We may want to include a value to ensure that non-locked stakers
    //       are unable to earn any additional ATokens; this should be reserved
    //       *only* for locked stakers. This may need to take the form of some
    //       variable inside Position; it is not completely clear.
    //
    //       Also, the exact types may be changed, but in no way show
    //       the change result in storing more than 32 additional bytes;
    //       this should not be necessary.
    struct Position {
        // number of aTokens when including weight
        uint224 weightedShares;
        // lockedPosition specifies if the staked position is locked;
        // this is only used for locked staking positions.
        bool lockedPosition;
        // number of aToken
        uint224 shares;
        // block number after which the position may be burned.
        // prevents double spend of voting weight
        uint32 freeAfter;
        // block number after which the position may be collected or burned.
        uint256 withdrawFreeAfter;
        // the last value of the ethState accumulator this account performed a
        // withdraw at
        uint256 accumulatorEth;
        // the last value of the tokenState accumulator this account performed a
        // withdraw at
        uint256 accumulatorToken;
    }

    // Accumulator is a struct that allows values to be collected such that the
    // remainders of floor division may be cleaned up
    struct Accumulator {
        // accumulator is a sum of all changes always increasing
        uint256 accumulator;
        // slush stores division remainders until they may be distributed evenly
        uint256 slush;
    }

    // Tier specifies the different levels of locking available in order
    // to allow for different risk tolerances.
    struct Tier {
        // multiplier specifies the weight associated with the staking tier;
        // see _LOCKING_TIER_DENOMINATOR for the denominator.
        uint256 multiplier;
        // lockDuration specifies the lock period for the Tier;
        // a Tier with larger multiplier comes with a longer locking period.
        uint32 lockDuration;
    }
    // Main Idea:
    //
    // There should be a range of possible tier levels.
    // The lowest tier (Tier 0) should not earn any rewards of the additional
    // ATokens which are minted; this should be reserved for higher tiers.
    // There will likely be a number of higher tiers for a locking period
    // between (say) 1 week and 12 months; the lengths will likely be determined
    // by epoch number (say 4 options: 1 week; 3 months; 6 months 12 months)
    //
    // Users should not be able to make their own Tier position;
    // rather, they will choose from the allowable options.

    // _MAX_MINT_LOCK describes the maximum interval a Position may be locked
    // during a call to mintTo
    uint256 internal constant _MAX_MINT_LOCK = 1051200;
    // 10**18
    uint256 internal constant _ACCUMULATOR_SCALE_FACTOR = 1000000000000000000;
    // constants for the cb state
    bool internal constant _CIRCUIT_BREAKER_OPENED = true;
    bool internal constant _CIRCUIT_BREAKER_CLOSED = false;

    // monotonically increasing counter
    uint256 internal _counter;

    // cb is the circuit breaker
    // cb is a set only object
    bool internal _circuitBreaker;

    // TODO: it is not clear if _shares and _reserveToken
    //       serve different roles. They appear to be the same.

    // _shares stores total amount of AToken staked in contract.
    // If this is used, we may want to rename this as _weightedShares.
    uint256 internal _shares;

    // _tokenState tracks distribution of AToken that originate from slashing
    // events
    Accumulator internal _tokenState;

    // _ethState tracks the distribution of Eth that originate from the sale of
    // BTokens
    Accumulator internal _ethState;

    // _positions tracks all staked positions based on tokenID
    mapping(uint256 => Position) internal _positions;

    // state to keep track of the amount of Eth deposited and collected from the
    // contract
    uint256 internal _reserveEth;

    // state to keep track of the amount of ATokens deposited and collected
    // from the contract.
    //
    // As noted above, it is not clear if this duplicates _shares global variable.
    // If this does duplicate _shares, this could be kept and used to represent
    // the total ATokens in the smart contract: the *unweighted* value.
    uint256 internal _reserveToken;

    // state to keep track of the amount of ATokens to be distributed
    // from the contract to locked stakers
    uint256 internal _additionalToken;

    // denominator used when computing weighted stake
    uint256 internal constant _LOCKING_TIER_DENOMINATOR = 1000000;

    // TODO: need different tiers; potential for wanting a short locking period.
    //       will probably need to use an enum. We may want a specific struct to
    //       to use for Tier objects.

    // _REWARD_ERA specifies the number of epochs per reward era,
    // which determines the specific decay rate of additional ATokens which
    // are minted each snapshot (epoch).
    //
    // TODO: we should think about what value we want to choose
    uint32 internal constant _REWARD_ERA = 4400;

    // _ADDITIONAL_ATOKENS specifies the total number of additional ATokens
    // which will be minted.
    //
    // TODO: this value must be determined
    uint256 internal constant _ADDITIONAL_ATOKENS = 220000000000000000000000000;
}
