### 

后端数据库采用MongoDB 

### 数据库表字段设计

1. project party user 表   
  存储项目方用户相关信息
``` 
public_address  last_login_time project_names
地址            登录时间         关联项目名单（数组，用来关联project whitelist manager）

```
2. common user 表  
  存储普通用户信息

```
public_address  last_login_time  project_names
地址             更新时间         关联项目名单（数组，用来关联project whitelist manager
```
3. project whitelist manager 表 （这里想法是一个项目建一张表，项目方新建一个项目就新建一个表） 
  存储具体项目的白名单信息
```
project_name project_address owner_address  whitelist_address  update_time  status
项目名    项目地址   管理员地址  白名单地址  操作更新时间 审批状态（未审批，通过，驳回三种）
```
4. project chain info 表 管理项目上的chain的相关信息

```
project_name project_address  private_key owner_address   chain_type chain_id
项目名 项目地址   授权私钥 项目方地址 chain的种类 chain的id
```

5. monitor_logs表 监控项目logs的表 设计是一个项目一张表 字段根据the graph的返回字段设置也可根据eth的logs日志设计
todo

```
{
  blockHash: "0x7eaf6abe64592d10828e136635aa6be6f4d09da3bb5b9fddf87773ee152d657c",
  blockNumber: 4654718,
  contractAddress: null,
  cumulativeGasUsed: 52464,
  from: "0x076979a0b3c87334e5d72e3afcafaa80f7888cac",
  gasUsed: 52464,
  logs: [{
      address: "0x73c2a5b1a32fa8e33101a6ab119203f4417feae4",
      blockHash: "0x7eaf6abe64592d10828e136635aa6be6f4d09da3bb5b9fddf87773ee152d657c",
      blockNumber: 4654718,
      data: "0x0000000000000000000000000000000000000000000000056bc75e2d63100000",
      logIndex: 0,
      removed: false,
      topics: ["0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef", "0x000000000000000000000000076979a0b3c87334e5d72e3afcafaa80f7888cac", "0x000000000000000000000000cd9f286ba6a3d2df7885f4a2be267fc524d32bd3"],
      transactionHash: "0xe03fac05ff4dde83fc9267184fd8c08bd78599f950e817dbf7fa4a4d4d319ce2",
      transactionIndex: 0
  }],
  logsBloom: "0x20000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000080000008000000000400000000000000000000000000000000000000040000000000000000100000000000000000000000000010000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000200000000000000000000000200000002000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000040000000000000000400",
  status: "0x1",
  to: "0x73c2a5b1a32fa8e33101a6ab119203f4417feae4",
  transactionHash: "0xe03fac05ff4dde83fc9267184fd8c08bd78599f950e817dbf7fa4a4d4d319ce2",
  transactionIndex: 0
}
```

```

contract_address topic1 topic2 topic3 topic4 data tx_hash block_hash block_number block_time index tx_index

```



6. 资产跨链管理表
