// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

/// @notice This contract tries to touch as many EVM opcodes as possible to test backend support.
/// Not every opcode is directly accessible in Solidity, but this covers all accessible ones.
/// For Foundry/forge test, see `test/OpcodeCoverage.t.sol`.
contract OpcodeCoverage {
  // Arithmetic
  function opADD(uint256 a, uint256 b) public pure returns (uint256) {
    assembly {
      a := add(a, b)
    }
    return a;
  }

  function opMUL(uint256 a, uint256 b) public pure returns (uint256) {
    assembly {
      a := mul(a, b)
    }
    return a;
  }

  function opSUB(uint256 a, uint256 b) public pure returns (uint256) {
    assembly {
      a := sub(a, b)
    }
    return a;
  }

  function opDIV(uint256 a, uint256 b) public pure returns (uint256) {
    assembly {
      a := div(a, b)
    }
    return a;
  }

  function opSDIV(int256 a, int256 b) public pure returns (int256) {
    assembly {
      a := sdiv(a, b)
    }
    return a;
  }

  function opMOD(uint256 a, uint256 b) public pure returns (uint256) {
    assembly {
      a := mod(a, b)
    }
    return a;
  }

  function opSMOD(int256 a, int256 b) public pure returns (int256) {
    assembly {
      a := smod(a, b)
    }
    return a;
  }

  function opADDMOD(uint256 a, uint256 b, uint256 m) public pure returns (uint256) {
    assembly {
      a := addmod(a, b, m)
    }
    return a;
  }

  function opMULMOD(uint256 a, uint256 b, uint256 m) public pure returns (uint256) {
    assembly {
      a := mulmod(a, b, m)
    }
    return a;
  }

  function opEXP(uint256 a, uint256 b) public pure returns (uint256) {
    assembly {
      a := exp(a, b)
    }
    return a;
  }

  function opSIGNEXTEND(uint256 b, int256 x) public pure returns (int256 y) {
      assembly { y := signextend(b, x) }
  }

  // Comparison/Bitwise
  function opLT(uint256 a, uint256 b) public pure returns (bool r) {
    assembly {
      r := lt(a, b)
    }
  }

  function opGT(uint256 a, uint256 b) public pure returns (bool r) {
    assembly {
      r := gt(a, b)
    }
  }

  function opSLT(int256 a, int256 b) public pure returns (bool r) {
    assembly {
      r := slt(a, b)
    }
  }

  function opSGT(int256 a, int256 b) public pure returns (bool r) {
    assembly {
      r := sgt(a, b)
    }
  }

  function opEQ(uint256 a, uint256 b) public pure returns (bool r) {
    assembly {
      r := eq(a, b)
    }
  }

  function opISZERO(uint256 a) public pure returns (bool r) {
    assembly {
      r := iszero(a)
    }
  }

  function opAND(uint256 a, uint256 b) public pure returns (uint256) {
    assembly {
      a := and(a, b)
    }
    return a;
  }

  function opOR(uint256 a, uint256 b) public pure returns (uint256) {
    assembly {
      a := or(a, b)
    }
    return a;
  }

  function opXOR(uint256 a, uint256 b) public pure returns (uint256) {
    assembly {
      a := xor(a, b)
    }
    return a;
  }

  function opNOT(uint256 a) public pure returns (uint256) {
    assembly {
      a := not(a)
    }
    return a;
  }

  function opBYTE(uint256 n, bytes32 x) public pure returns (uint8 r) {
    assembly {
      r := byte(n, x)
    }
  }

  function opSHL(uint256 a, uint256 b) public pure returns (uint256) {
    assembly {
      a := shl(a, b)
    }
    return a;
  }

  function opSHR(uint256 shift, uint256 value) public pure returns (uint256 r) {
    assembly { r := shr(shift, value) }
}

  function opSAR(int256 a, int256 b) public pure returns (int256) {
    assembly {a := sar(a, b)}
    return a;
  }

  // Crypto
  function opKECCAK256(bytes memory d) public pure returns (bytes32 h) {
    assembly {
      h := keccak256(add(d, 32), mload(d))
    }
  }

  // Environmental/Info
  function opADDRESS() public view returns (address a) {
    assembly {
      a := address()
    }
  }

  function opBALANCE(address a) public view returns (uint256 b) {
    assembly {
      b := balance(a)
    }
  }

  function opORIGIN() public view returns (address o) {
    assembly {
      o := origin()
    }
  }

  function opCALLER() public view returns (address c) {
    assembly {
      c := caller()
    }
  }

  function opCALLVALUE() public payable returns (uint256 v) {
    assembly {
      v := callvalue()
    }
  }

  function opCALLDATALOAD(uint256 i) public pure returns (uint256 v) {
    assembly {
      v := calldataload(i)
    }
  }

  function opCALLDATASIZE() public pure returns (uint256 s) {
    assembly {
      s := calldatasize()
    }
  }

  function opCODESIZE() public pure returns (uint256 s) {
    assembly {
      s := codesize()
    }
  }

  function opGASPRICE() public view returns (uint256 p) {
    assembly {
      p := gasprice()
    }
  }

  function opCHAINID() public view returns (uint256 id) {
    assembly {
      id := chainid()
    }
  }

  function opSELFBALANCE() public view returns (uint256 v) {
    assembly {
      v := selfbalance()
    }
  }

  function opBASEFEE() public view returns (uint256 v) {
    assembly {
      v := basefee()
    }
  }

  // Block Info
  function opBLOCKHASH(uint256 n) public view returns (bytes32 r) {
      assembly { r := blockhash(n) }
  }

  function opCOINBASE() public view returns (address a) {
    assembly {
      a := coinbase()
    }
  }

  function opTIMESTAMP() public view returns (uint256 t) {
    assembly {
      t := timestamp()
    }
  }

  function opNUMBER() public view returns (uint256 n) {
    assembly {
      n := number()
    }
  }

  function opDIFFICULTY() public view returns (uint256 d) {
    assembly {
      d := prevrandao()
    }
  }

  function opGASLIMIT() public view returns (uint256 l) {
    assembly {
      l := gaslimit()
    }
  }

  // Storage/Memory
  function opSSTORE(bytes32 k, bytes32 v) public {
    assembly {
      sstore(k, v)
    }
  }

  function opSLOAD(bytes32 k) public view returns (bytes32 v) {
    assembly {
      v := sload(k)
    }
  }

  function opMSTORE(uint256 p, uint256 v) public pure returns (uint256 o) {
    assembly {
      mstore(p, v)
      o := mload(p)
    }
  }

  function opMSTORE8(uint256 p, uint8 v) public pure returns (uint8 o) {
    assembly {mstore8(p, v) o := mload(p)}
  }

  function opMLOAD(uint256 p) public pure returns (uint256 v) {
    assembly {
      v := mload(p)
    }
  }

  // Logging
  event Log0(bytes32 a);

  function opLOG0(bytes32 a) public {
    assembly {
      log0(a, 0x20)
    }
    emit Log0(a);
  }

  // Misc
  function opGAS() public view returns (uint256 g) {
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
  // function opPC() public view returns (uint256 p) {
  //     assembly { p := pc() }
  // }
  // function opMSIZE() public pure returns (uint256 s) {
  //     assembly { s := msize() }
  // }

  // Control flow (not accessible directly: JUMP, JUMPI, JUMPDEST, STOP, REVERT, etc.)

  // Pushes, Dups, Swaps are part of bytecode, not accessible directly in Solidity.

  // CREATE2, CREATE, CALL, STATICCALL, DELEGATECALL, SELFDESTRUCT are possible but require more code
  // Add as needed.
}
