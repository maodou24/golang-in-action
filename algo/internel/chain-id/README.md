# Chain Id

系统中针对网络连接的设备，设计了一种Chain ID的结构，每个网络设备可以设置Chain ID，如果一个设备通过另一个设备接入网络，需要基于副设备的Chain ID设置自己的Chain ID

举例：
```tree
device1: 1
    device2: 1.1
    device3: 1.2
        device5: 1.2.2
    device5: 1.4.2
device4: 2
```

设计一个算法结构：
1. 同一链上的从链尾往链首遍历依次工作，无Chain ID链接的设备可以并行工作（device2和device3）
2. 即使中间部分设备节点不存在，要找到最终的父节点，如device5要在device1之前

Solution：
前缀树



