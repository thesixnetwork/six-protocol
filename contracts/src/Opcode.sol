// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

/// @notice This contract tries to touch as many EVM opcodes as possible to test backend support.
/// Not every opcode is directly accessible in Solidity, but this covers all accessible ones.
/// For Foundry/forge test, see `test/OpcodeCoverage.t.sol`.
contract OpcodeCoverage {
  // Arithmetic
  function op_ADD(uint256 a, uint256 b) public pure returns (uint256) {
    assembly {
      a := add(a, b)
    }
    return a;
  }

  function op_MUL(uint256 a, uint256 b) public pure returns (uint256) {
    assembly {
      a := mul(a, b)
    }
    return a;
  }

  function op_SUB(uint256 a, uint256 b) public pure returns (uint256) {
    assembly {
      a := sub(a, b)
    }
    return a;
  }

  function op_DIV(uint256 a, uint256 b) public pure returns (uint256) {
    assembly {
      a := div(a, b)
    }
    return a;
  }

  function op_SDIV(int256 a, int256 b) public pure returns (int256) {
    assembly {
      a := sdiv(a, b)
    }
    return a;
  }

  function op_MOD(uint256 a, uint256 b) public pure returns (uint256) {
    assembly {
      a := mod(a, b)
    }
    return a;
  }

  function op_SMOD(int256 a, int256 b) public pure returns (int256) {
    assembly {
      a := smod(a, b)
    }
    return a;
  }

  function op_ADDMOD(uint256 a, uint256 b, uint256 m) public pure returns (uint256) {
    assembly {
      a := addmod(a, b, m)
    }
    return a;
  }

  function op_MULMOD(uint256 a, uint256 b, uint256 m) public pure returns (uint256) {
    assembly {
      a := mulmod(a, b, m)
    }
    return a;
  }

  function op_EXP(uint256 a, uint256 b) public pure returns (uint256) {
    assembly {
      a := exp(a, b)
    }
    return a;
  }

  function op_SIGNEXTEND(uint256 b, int256 x) public pure returns (int256 y) {
      assembly { y := signextend(b, x) }
  }

  // Comparison/Bitwise
  function op_LT(uint256 a, uint256 b) public pure returns (bool r) {
    assembly {
      r := lt(a, b)
    }
  }

  function op_GT(uint256 a, uint256 b) public pure returns (bool r) {
    assembly {
      r := gt(a, b)
    }
  }

  function op_SLT(int256 a, int256 b) public pure returns (bool r) {
    assembly {
      r := slt(a, b)
    }
  }

  function op_SGT(int256 a, int256 b) public pure returns (bool r) {
    assembly {
      r := sgt(a, b)
    }
  }

  function op_EQ(uint256 a, uint256 b) public pure returns (bool r) {
    assembly {
      r := eq(a, b)
    }
  }

  function op_ISZERO(uint256 a) public pure returns (bool r) {
    assembly {
      r := iszero(a)
    }
  }

  function op_AND(uint256 a, uint256 b) public pure returns (uint256) {
    assembly {
      a := and(a, b)
    }
    return a;
  }

  function op_OR(uint256 a, uint256 b) public pure returns (uint256) {
    assembly {
      a := or(a, b)
    }
    return a;
  }

  function op_XOR(uint256 a, uint256 b) public pure returns (uint256) {
    assembly {
      a := xor(a, b)
    }
    return a;
  }

  function op_NOT(uint256 a) public pure returns (uint256) {
    assembly {
      a := not(a)
    }
    return a;
  }

  function op_BYTE(uint256 n, bytes32 x) public pure returns (uint8 r) {
    assembly {
      r := byte(n, x)
    }
  }

  function op_SHL(uint256 a, uint256 b) public pure returns (uint256) {
    assembly {
      a := shl(a, b)
    }
    return a;
  }

  function op_SHR(uint256 shift, uint256 value) public pure returns (uint256 r) {
    assembly { r := shr(shift, value) }
}

  function op_SAR(int256 a, int256 b) public pure returns (int256) {
    assembly {a := sar(a, b)}
    return a;
  }

  // Crypto
  function op_KECCAK256(bytes memory d) public pure returns (bytes32 h) {
    assembly {
      h := keccak256(add(d, 32), mload(d))
    }
  }

  // Environmental/Info
  function op_ADDRESS() public view returns (address a) {
    assembly {
      a := address()
    }
  }

  function op_BALANCE(address a) public view returns (uint256 b) {
    assembly {
      b := balance(a)
    }
  }

  function op_ORIGIN() public view returns (address o) {
    assembly {
      o := origin()
    }
  }

  function op_CALLER() public view returns (address c) {
    assembly {
      c := caller()
    }
  }

  function op_CALLVALUE() public payable returns (uint256 v) {
    assembly {
      v := callvalue()
    }
  }

  function op_CALLDATALOAD(uint256 i) public pure returns (uint256 v) {
    assembly {
      v := calldataload(i)
    }
  }

  function op_CALLDATASIZE() public pure returns (uint256 s) {
    assembly {
      s := calldatasize()
    }
  }

  function op_CODESIZE() public pure returns (uint256 s) {
    assembly {
      s := codesize()
    }
  }

  function op_GASPRICE() public view returns (uint256 p) {
    assembly {
      p := gasprice()
    }
  }

  function op_CHAINID() public view returns (uint256 id) {
    assembly {
      id := chainid()
    }
  }

  function op_SELFBALANCE() public view returns (uint256 v) {
    assembly {
      v := selfbalance()
    }
  }

  function op_BASEFEE() public view returns (uint256 v) {
    assembly {
      v := basefee()
    }
  }

  // Block Info
  function op_BLOCKHASH(uint256 n) public view returns (bytes32 r) {
      assembly { r := blockhash(n) }
  }

  function op_COINBASE() public view returns (address a) {
    assembly {
      a := coinbase()
    }
  }

  function op_TIMESTAMP() public view returns (uint256 t) {
    assembly {
      t := timestamp()
    }
  }

  function op_NUMBER() public view returns (uint256 n) {
    assembly {
      n := number()
    }
  }

  function op_DIFFICULTY() public view returns (uint256 d) {
    assembly {
      d := prevrandao()
    }
  }

  function op_GASLIMIT() public view returns (uint256 l) {
    assembly {
      l := gaslimit()
    }
  }

  // Storage/Memory
  function op_SSTORE(bytes32 k, bytes32 v) public {
    assembly {
      sstore(k, v)
    }
  }

  function op_SLOAD(bytes32 k) public view returns (bytes32 v) {
    assembly {
      v := sload(k)
    }
  }

  function op_MSTORE(uint256 p, uint256 v) public pure returns (uint256 o) {
    assembly {
      mstore(p, v)
      o := mload(p)
    }
  }

  function op_MSTORE8(uint256 p, uint8 v) public pure returns (uint8 o) {
    assembly {mstore8(p, v) o := mload(p)}
  }

  function op_MLOAD(uint256 p) public pure returns (uint256 v) {
    assembly {
      v := mload(p)
    }
  }

  // Logging
  event Log0(bytes32 a);

  function op_LOG0(bytes32 a) public {
    assembly {
      log0(a, 0x20)
    }
    emit Log0(a);
  }

  // Misc
  function op_GAS() public view returns (uint256 g) {
    assembly {
      g := gas()
    }
  }

  function testSignExtend() public pure returns (int256) {
    int256 x;
    assembly {
      x := signextend(1, 0x80)
    }
    return x;
  }
  // function op_PC() public view returns (uint256 p) {
  //     assembly { p := pc() }
  // }
  // function op_MSIZE() public pure returns (uint256 s) {
  //     assembly { s := msize() }
  // }

  // Control flow (not accessible directly: JUMP, JUMPI, JUMPDEST, STOP, REVERT, etc.)

  // Pushes, Dups, Swaps are part of bytecode, not accessible directly in Solidity.

  // CREATE2, CREATE, CALL, STATICCALL, DELEGATECALL, SELFDESTRUCT are possible but require more code
  // Add as needed.
}
