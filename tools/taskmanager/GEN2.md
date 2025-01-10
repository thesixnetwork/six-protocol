# GEN2 Task Management

---

## CLOSED

| **Date**   | **Task**                |
| ---------- | ----------------------- |
| 2024-xx-xx | EVM Precompile contract |

---

## OPEN

| **Date**   | **Task**                                                           | **Priority** | **Owner** | **Tags**                   |
| ---------- | ------------------------------------------------------------------ | ------------ | --------- | -------------------------- |
| 2025-12-24 | Remove NFT schema info redundancy                                  | Low          | @dDeedev  | `#schema`, `#storage`      |
| 2025-12-24 | Query value from precompile and reuse with other solidity contract | Medium       | @dDeedev  | `#precompile`, `#solidity` |
| 2025-12-24 | Revive NFT Oracle                                                  | Medium       | @dDeedev  | `#oracle`                  |

---

## IN PROGRESS

| **Open Date** | **Task**       | **Priority** | Target Completion | **Owner** | **Tags** |
| ------------- | -------------- | ------------ | ----------------- | --------- | -------- |
| 2025-12-24    | Virtual Schema | High         | 2025-01-30        | @dDeedev  | `#gen2`  |

### Virtual Schema Tasks

| **ID** | **Task**                                                               | **Target Completion** | **Completed** | **Notes**   |
| ------ | ---------------------------------------------------------------------- | --------------------- | ------------- | ----------- |
| 1      | Implement Message Vote Virtual Schema                                  | 2024-12-31            | 2024-12-26    | Done        |
| 2      | Proposal Process                                                       | 2024-12-31            | 2024-12-26    | Done        |
| 3      | Action Policy                                                          | 2024-12-31            | 2024-12-26    | Done        |
| 4      | Enable and Disable Virtual Action                                      | 2024-12-31            | 2024-12-26    | Done        |
| 5      | Make engine able to read value across schema                           | 2024-12-31            | 2024-12-26    | Done        |
| 6      | Hook Proposal height has met                                           | 2024-12-31            | 2024-12-26    | Done        |
| 7      | CrossSchema Struct                                                     | 2024-12-31            | 2024-12-26    | Done        |
| 8      | Implement Metadata's Method to CrossSchema                             | 2024-12-31            | 2024-12-26    | Done        |
| 9      | VirtualSchema required token id of each src schema                     | 2024-12-31            | 2024-12-26    | Done        |
| 10     | Each src schema must contain info of virtual action                    | 2024-12-31            | 2024-12-26    | Done        |
| 11     | End of SubmitTime hook                                                 | 2024-12-31            | 2024-12-26    | Done        |
| 12     | Edit Virtual Action                                                    | 2024-12-31            | 2024-12-26    | Done        |
| 13     | [BUG] Add Duplicate Virtual Action on the same Schema                  | 2025-01-10            | 2025-01-10    | Fixed       |
| 14     | [feat] Precompile Virtual Action                                       | 2025-01-15            | TBD           | In Progress |
| 15     | [chore] Create virtual schema must contain virtual action (on request) | 2025-01-15            | 2025-01-03    | In Progress |
| 16     | [chore] Combine enable/disable to change virtual schema                | 2025-01-15            | TBD           | Pending     |
| 17     | [chore] Lock virtual schema name                                       | 2025-01-15            | TBD           | Pending     |

---

## REVIEW

| **Date**   | **Task**                                                                                             | **Status** |
| ---------- | ---------------------------------------------------------------------------------------------------- | ---------- |
| 2024-12-24 | Action by admin does not require ref id                                                              | Reviewed   |
| 2025-01-08 | Test Migrate module from `github.com/thesixnetwork/sixnft` to `github.com/thesixnetwork/sixprotocol` | Reviewed   |
