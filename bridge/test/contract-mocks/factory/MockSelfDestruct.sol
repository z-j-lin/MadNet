// SPDX-License-Identifier: MIT-open-group
pragma solidity ^0.8.11;

import "contracts/libraries/proxy/ProxyInternalUpgradeLock.sol";
import "contracts/libraries/proxy/ProxyInternalUpgradeUnlock.sol";

/// @custom:salt MockSelfDestruct
contract MockSelfDestruct is ProxyInternalUpgradeLock, ProxyInternalUpgradeUnlock {
    address immutable _factory;
    uint256 public v;
    uint256 public immutable i;

    constructor(uint256 _i, bytes memory) {
        i = _i;
        _factory = msg.sender;
    }

    function getFactory() external view returns (address) {
        return _factory;
    }
}
