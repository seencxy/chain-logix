// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

// 主合约定义
contract LogisticsManagement {
    // 定义Package结构体
    struct Package {
        string packageId; // 包裹ID
        string sender; // 发送者用户名
        string receiver; // 接收者用户名
        string sender_address; // 发送者用户地址
        string receiver_address; // 接受者地址
        string current_address; // 当前所在地址
        string supply_chain_admin; // 物流管理人员
        string status; // 包裹当前状态 await 等待接收 ship 运输中 service 送达
    }

    // 合约所有者的地址
    address public owner;

    // 存储状态为 'await' 的包裹ID
    string[] private awaitPackages;
    string[] private allPackageIds;

    // 包裹ID到Package映射
    mapping(string => Package) public packages;

    // 事件定义
    event PackageStatusUpdated(
        string packageId,
        string username,
        string newStatus
    );

    // 构造函数，设置合约部署者为所有者
    constructor() {
        owner = msg.sender;
    }

    // 注册新包裹的函数
    function registerPackage(
        string memory _packageId,
        string memory _sender,
        string memory _receiver,
        string memory _senderAddress,
        string memory _receiverAddress
    ) public {
        // 创建新包裹并存入映射
        packages[_packageId] = Package(
            _packageId,
            _sender,
            _receiver,
            _senderAddress,
            _receiverAddress,
            _senderAddress,
            unicode"", // 等待物流管理人员接单
            "await"
        );

        awaitPackages.push(_packageId); // 将新包裹ID添加到等待列表
        allPackageIds.push(_packageId);
    }

    // 物流管理员接收包裹
    function supplyAdminAcceptPackage(
        string memory _adminName,
        string memory _packageId
    ) public {
        // 更新物流状态
        packages[_packageId].status = "ship";
        packages[_packageId].supply_chain_admin = _adminName;
        // 触发事件
        emit PackageStatusUpdated(
            _packageId,
            packages[_packageId].sender,
            packages[_packageId].status
        );

        removeAwaitPackage(_packageId);
    }

    // 更新包裹当前地址的函数
    function updateCurrentAddress(
        string memory _adminName,
        string memory _packageId,
        string memory _newAddress
    ) public {
        require(
            bytes(packages[_packageId].sender_address).length > 0,
            "Package does not exist."
        );

        require(
            keccak256(
                abi.encodePacked(packages[_packageId].supply_chain_admin)
            ) == keccak256(abi.encodePacked(_adminName)),
            "this admin not accept admin"
        );

        packages[_packageId].current_address = _newAddress;

        // 如果新的地址和接受地址相同则更新物流状况为完成
        if (
            keccak256(abi.encodePacked(_newAddress)) ==
            keccak256(abi.encodePacked(packages[_packageId].receiver_address))
        ) {
            // 更新状态将状态设置为送达
            packages[_packageId].status = "service";
        }

        // 触发包裹地址更新事件，包含更新者的用户名和包裹ID
        emit PackageStatusUpdated(
            _packageId,
            packages[_packageId].sender,
            packages[_packageId].status
        );
    }

    // 获取包裹详情的函数
    function getPackage(string memory _packageId)
    public
    view
    returns (Package memory)
    {
        require(
            bytes(packages[_packageId].sender_address).length > 0,
            "Package does not exist."
        );
        // 返回包裹信息
        return packages[_packageId];
    }

    // 获取所有状态为 'await' 的包裹
    function getAllAwaitPackages() public view returns (Package[] memory) {
        Package[] memory packagesInAwait = new Package[](awaitPackages.length);
        for (uint256 i = 0; i < awaitPackages.length; i++) {
            packagesInAwait[i] = packages[awaitPackages[i]];
        }
        return packagesInAwait;
    }

    // 从等待列表中移除特定包裹ID的辅助函数
    function removeAwaitPackage(string memory _packageId) private {
        for (uint256 i = 0; i < awaitPackages.length; i++) {
            if (
                keccak256(abi.encodePacked(awaitPackages[i])) ==
                keccak256(abi.encodePacked(_packageId))
            ) {
                awaitPackages[i] = awaitPackages[awaitPackages.length - 1];
                awaitPackages.pop();
                break;
            }
        }
    }

    // 获取所有包裹的信息
    function getAllPackages() public view returns (Package[] memory) {
        Package[] memory allPackages = new Package[](allPackageIds.length);
        for (uint256 i = 0; i < allPackageIds.length; i++) {
            allPackages[i] = packages[allPackageIds[i]];
        }
        return allPackages;
    }

    // 获取指定用户名的包裹
    function getPackagesByUsername(string memory username)
    public
    view
    returns (Package[] memory)
    {
        uint256 count = 0;

        // 首先，计算与用户名相关的包裹数量
        for (uint256 i = 0; i < allPackageIds.length; i++) {
            if (
                keccak256(
                    abi.encodePacked(packages[allPackageIds[i]].sender)
                ) ==
                keccak256(abi.encodePacked(username)) ||
                keccak256(
                    abi.encodePacked(packages[allPackageIds[i]].receiver)
                ) ==
                keccak256(abi.encodePacked(username))
            ) {
                count++;
            }
        }

        // 创建一个适当大小的数组来存储这些包裹
        Package[] memory userPackages = new Package[](count);
        uint256 index = 0;

        // 再次遍历，将与用户名相关的包裹添加到数组中
        for (uint256 i = 0; i < allPackageIds.length; i++) {
            if (
                keccak256(
                    abi.encodePacked(packages[allPackageIds[i]].sender)
                ) ==
                keccak256(abi.encodePacked(username)) ||
                keccak256(
                    abi.encodePacked(packages[allPackageIds[i]].receiver)
                ) ==
                keccak256(abi.encodePacked(username))
            ) {
                userPackages[index] = packages[allPackageIds[i]];
                index++;
            }
        }

        return userPackages;
    }

    // 获取指定物流管理员负责的所有包裹
    function getPackagesByAdmin(string memory _adminName)
    public
    view
    returns (Package[] memory)
    {
        uint256 count = 0;

        // 计算该管理员负责的包裹数量
        for (uint256 i = 0; i < allPackageIds.length; i++) {
            if (
                keccak256(
                    abi.encodePacked(
                        packages[allPackageIds[i]].supply_chain_admin
                    )
                ) == keccak256(abi.encodePacked(_adminName))
            ) {
                count++;
            }
        }

        // 创建一个数组存储这些包裹
        Package[] memory adminPackages = new Package[](count);
        uint256 index = 0;

        // 再次遍历，添加该管理员负责的包裹到数组
        for (uint256 i = 0; i < allPackageIds.length; i++) {
            if (
                keccak256(
                    abi.encodePacked(
                        packages[allPackageIds[i]].supply_chain_admin
                    )
                ) == keccak256(abi.encodePacked(_adminName))
            ) {
                adminPackages[index] = packages[allPackageIds[i]];
                index++;
            }
        }

        return adminPackages;
    }

    // 取消状态为 'await' 的包裹
    function cancelAwaitPackage(string memory _packageId) public {
        // 确保包裹存在且状态为 'await'
        require(
            bytes(packages[_packageId].packageId).length > 0 &&
            keccak256(abi.encodePacked(packages[_packageId].status)) ==
            keccak256(abi.encodePacked("await")),
            "Package must exist and be in 'await' status."
        );

        // 设置包裹状态为 'cancelled'
        packages[_packageId].status = "cancelled";

        // 从等待列表中移除包裹
        removeAwaitPackage(_packageId);

        // 触发包裹状态更新事件
        emit PackageStatusUpdated(
            _packageId,
            packages[_packageId].sender,
            "cancelled"
        );
    }
}
