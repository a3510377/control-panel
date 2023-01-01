# permissions

| `permission`    | `value`            | `description` | `type` |
| --------------- | ------------------ | ------------- | :----: |
| `None`          | `0x0000`           |               | `S,I`  |
| `Administrator` | `0x0001 (1 << 00)` |               | `S,I`  |
| `ManageUser`    | `0x0002 (1 << 01)` |               | `S,I`  |
| `ManageNick`    | `0x0004 (1 << 02)` |               | `S,I`  |
| `ChangeNick`    | `0x0008 (1 << 03)` |               | `S,I`  |
| `ManageFile`    | `0x0010 (1 << 04)` |               |  `I`   |
| `ManageEnv`     | `0x0020 (1 << 05)` |               |  `I`   |
| `ReadEnv`       | `0x0040 (1 << 06)` |               |  `I`   |
| `ReadFile`      | `0x0080 (1 << 07)` |               |  `I`   |
| `ChangeState`   | `0x0100 (1 << 08)` |               |  `I`   |
| `SendCommand`   | `0x0200 (1 << 09)` |               |  `I`   |
| `ReadState`     | `0x0400 (1 << 10)` |               | `S,I`  |
| `ViewLog`       | `0x0800 (1 << 11)` |               | `S,I`  |
