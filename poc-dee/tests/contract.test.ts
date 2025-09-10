import {
  assert,
  describe,
  test,
  clearStore,
  beforeAll,
  afterAll
} from "matchstick-as/assembly/index"
import { Address, BigInt } from "@graphprotocol/graph-ts"
import { AllAssetUpdatersRevoked } from "../generated/schema"
import { AllAssetUpdatersRevoked as AllAssetUpdatersRevokedEvent } from "../generated/Contract/Contract"
import { handleAllAssetUpdatersRevoked } from "../src/contract"
import { createAllAssetUpdatersRevokedEvent } from "./contract-utils"

// Tests structure (matchstick-as >=0.5.0)
// https://thegraph.com/docs/en/subgraphs/developing/creating/unit-testing-framework/#tests-structure

describe("Describe entity assertions", () => {
  beforeAll(() => {
    let newAllAssetUpdatersRevokedEvent = createAllAssetUpdatersRevokedEvent()
    handleAllAssetUpdatersRevoked(newAllAssetUpdatersRevokedEvent)
  })

  afterAll(() => {
    clearStore()
  })

  // For more test scenarios, see:
  // https://thegraph.com/docs/en/subgraphs/developing/creating/unit-testing-framework/#write-a-unit-test

  test("AllAssetUpdatersRevoked created and stored", () => {
    assert.entityCount("AllAssetUpdatersRevoked", 1)

    // 0xa16081f360e3847006db660bae1c6d1b2e17ec2a is the default address used in newMockEvent() function

    // More assert options:
    // https://thegraph.com/docs/en/subgraphs/developing/creating/unit-testing-framework/#asserts
  })
})
