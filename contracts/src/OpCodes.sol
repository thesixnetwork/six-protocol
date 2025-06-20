// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

/// @notice This contract tries to touch as many EVM opcodes as possible to test backend support.
/// Not every opcode is directly accessible in Solidity, but this covers all accessible ones.
/// For Foundry/forge test, see `test/OpcodeCoverage.t.sol`.
contract OpcodeCoverage {
  // Arithmetic
  function opAdd(uint256 a, uint256 b) public pure returns (uint256) {
    assembly {
      a := add(a, b)
    }
    return a;
  }

  function opMul(uint256 a, uint256 b) public pure returns (uint256) {
    assembly {
      a := mul(a, b)
    }
    return a;
  }

  function opSub(uint256 a, uint256 b) public pure returns (uint256) {
    assembly {
      a := sub(a, b)
    }
    return a;
  }

  function opDiv(uint256 a, uint256 b) public pure returns (uint256) {
    assembly {
      a := div(a, b)
    }
    return a;
  }

  function opSdiv(int256 a, int256 b) public pure returns (int256) {
    assembly {
      a := sdiv(a, b)
    }
    return a;
  }

  function opMod(uint256 a, uint256 b) public pure returns (uint256) {
    assembly {
      a := mod(a, b)
    }
    return a;
  }

  function opSmod(int256 a, int256 b) public pure returns (int256) {
    assembly {
      a := smod(a, b)
    }
    return a;
  }

  function opAddmod(uint256 a, uint256 b, uint256 m) public pure returns (uint256) {
    assembly {
      a := addmod(a, b, m)
    }
    return a;
  }

  function opMulmod(uint256 a, uint256 b, uint256 m) public pure returns (uint256) {
    assembly {
      a := mulmod(a, b, m)
    }
    return a;
  }

  function opExp(uint256 a, uint256 b) public pure returns (uint256) {
    assembly {
      a := exp(a, b)
    }
    return a;
  }

  function opSignextend(uint256 b, int256 x) public pure returns (int256 y) {
      assembly { y := signextend(b, x) }
  }

  // Comparison/Bitwise
  function opLt(uint256 a, uint256 b) public pure returns (bool r) {
    assembly {
      r := lt(a, b)
    }
  }

  function opGt(uint256 a, uint256 b) public pure returns (bool r) {
    assembly {
      r := gt(a, b)
    }
  }

  function opSlt(int256 a, int256 b) public pure returns (bool r) {
    assembly {
      r := slt(a, b)
    }
  }

  function opSgt(int256 a, int256 b) public pure returns (bool r) {
    assembly {
      r := sgt(a, b)
    }
  }

  function opEq(uint256 a, uint256 b) public pure returns (bool r) {
    assembly {
      r := eq(a, b)
    }
  }

  function opIszero(uint256 a) public pure returns (bool r) {
    assembly {
      r := iszero(a)
    }
  }

  function opAnd(uint256 a, uint256 b) public pure returns (uint256) {
    assembly {
      a := and(a, b)
    }
    return a;
  }

  function opOr(uint256 a, uint256 b) public pure returns (uint256) {
    assembly {
      a := or(a, b)
    }
    return a;
  }

  function opXor(uint256 a, uint256 b) public pure returns (uint256) {
    assembly {
      a := xor(a, b)
    }
    return a;
  }

  function opNot(uint256 a) public pure returns (uint256) {
    assembly {
      a := not(a)
    }
    return a;
  }

  function opByte(uint256 n, bytes32 x) public pure returns (uint8 r) {
    assembly {
      r := byte(n, x)
    }
  }

  function opShl(uint256 a, uint256 b) public pure returns (uint256) {
    assembly {
      a := shl(a, b)
    }
    return a;
  }

  function opShr(uint256 shift, uint256 value) public pure returns (uint256 r) {
    assembly { r := shr(shift, value) }
}

  function opSar(int256 a, int256 b) public pure returns (int256) {
    assembly {a := sar(a, b)}
    return a;
  }

  // Crypto
  function opKeccak256(bytes memory d) public pure returns (bytes32 h) {
    assembly {
      h := keccak256(add(d, 32), mload(d))
    }
  }
  
  // More gas-efficient hash function for calldata input
  function hashCalldata(bytes calldata data) public pure returns (bytes32 result) {
    assembly {
      // Get the calldata pointer directly
      let ptr := data.offset
      // Get the length of calldata bytes
      let len := data.length
      // Compute hash directly from calldata (saves gas by avoiding memory copies)
      result := keccak256(ptr, len)
    }
  }
  
  // Gas-efficient hash function for multiple inputs
  function hashPacked(uint256 a, address b, bytes32 c) public pure returns (bytes32 result) {
    assembly {
      // Allocate memory for the packed data (32 + 20 + 32 = 84 bytes)
      let ptr := mload(0x40) // Get free memory pointer
      
      // Store the values tightly packed
      mstore(ptr, a)               // uint256 uses 32 bytes
      mstore(add(ptr, 32), b)      // address padded to 32 bytes
      mstore(add(ptr, 64), c)      // bytes32 uses 32 bytes
      
      // Compute hash
      result := keccak256(ptr, 96)
      
      // No need to update the free memory pointer as we're not allocating permanently
    }
  }

  // Environmental/Info
  function opAddress() public view returns (address a) {
    assembly {
      a := address()
    }
  }

  function opBalance(address a) public view returns (uint256 b) {
    assembly {
      b := balance(a)
    }
  }

  function opOrigin() public view returns (address o) {
    assembly {
      o := origin()
    }
  }

  function opCaller() public view returns (address c) {
    assembly {
      c := caller()
    }
  }

  function opCallvalue() public payable returns (uint256 v) {
    assembly {
      v := callvalue()
    }
  }

  function opCalldataload(uint256 i) public pure returns (uint256 v) {
    assembly {
      v := calldataload(i)
    }
  }

  function opCalldatasize() public pure returns (uint256 s) {
    assembly {
      s := calldatasize()
    }
  }

  function opCodesize() public pure returns (uint256 s) {
    assembly {
      s := codesize()
    }
  }

  function opGasprice() public view returns (uint256 p) {
    assembly {
      p := gasprice()
    }
  }

  function opChainid() public view returns (uint256 id) {
    assembly {
      id := chainid()
    }
  }

  function opSelfbalance() public view returns (uint256 v) {
    assembly {
      v := selfbalance()
    }
  }

  function opBasefee() public view returns (uint256 v) {
    assembly {
      v := basefee()
    }
  }

  // Block Info
  function opBlockhash(uint256 n) public view returns (bytes32 r) {
      assembly { r := blockhash(n) }
  }

  function opCoinbase() public view returns (address a) {
    assembly {
      a := coinbase()
    }
  }

  function opTimestamp() public view returns (uint256 t) {
    assembly {
      t := timestamp()
    }
  }

  function opNumber() public view returns (uint256 n) {
    assembly {
      n := number()
    }
  }

  function opDifficulty() public view returns (uint256 d) {
    assembly {
      d := prevrandao()
    }
  }

  function opGaslimit() public view returns (uint256 l) {
    assembly {
      l := gaslimit()
    }
  }

  // Storage/Memory
  function opSstore(bytes32 k, bytes32 v) public {
    assembly {
      sstore(k, v)
    }
  }

  function opSload(bytes32 k) public view returns (bytes32 v) {
    assembly {
      v := sload(k)
    }
  }

  function opMstore(uint256 p, uint256 v) public pure returns (uint256 o) {
    assembly {
      mstore(p, v)
      o := mload(p)
    }
  }

  function opMstore8(uint256 p, uint8 v) public pure returns (uint8 o) {
    assembly {mstore8(p, v) o := mload(p)}
  }

  function opMload(uint256 p) public pure returns (uint256 v) {
    assembly {
      v := mload(p)
    }
  }

  // Logging
  event Log0(bytes32 a);

  function opLog0(bytes32 a) public {
    assembly {
      log0(a, 0x20)
    }
    emit Log0(a);
  }

  // Misc
  function opGas() public view returns (uint256 g) {
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
  // function opPc() public view returns (uint256 p) {
  //     assembly { p := pc() }
  // }
  // function opMsize() public pure returns (uint256 s) {
  //     assembly { s := msize() }
  // }

  // Control flow (not accessible directly: JUMP, JUMPI, JUMPDEST, STOP, REVERT, etc.)

  // Pushes, Dups, Swaps are part of bytecode, not accessible directly in Solidity.

  // CREATE2, CREATE, CALL, STATICCALL, DELEGATECALL, SELFDESTRUCT are possible but require more code
  // Add as needed.
}
