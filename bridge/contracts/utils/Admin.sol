// SPDX-License-Identifier: MIT-open-group
pragma solidity ^0.8.0;
import {AdminErrorCodes} from "contracts/libraries/errorCodes/AdminErrorCodes.sol";

abstract contract Admin {
    // _admin is a privileged role
    address internal _admin;

    /// @dev onlyAdmin enforces msg.sender is _admin
    modifier onlyAdmin() {
        require(
            msg.sender == _admin,
            string(abi.encodePacked(AdminErrorCodes.ADMIN_SENDER_MUST_BE_ADMIN))
        );
        _;
    }

    constructor(address admin_) {
        _admin = admin_;
    }

    /// @dev assigns a new admin may only be called by _admin
    function setAdmin(address admin_) public virtual onlyAdmin {
        _setAdmin(admin_);
    }

    /// @dev getAdmin returns the current _admin
    function getAdmin() public view returns (address) {
        return _admin;
    }

    // assigns a new admin may only be called by _admin
    function _setAdmin(address admin_) internal {
        _admin = admin_;
    }
}
