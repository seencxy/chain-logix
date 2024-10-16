// SPDX-License-Identifier: MIT  
pragma solidity ^0.8.0; 
  
contract DataUploader {  
    mapping(uint256 => string) private strings;  
    uint256 private stringCount;  
  
    event StringStored(uint256 indexed id, string data);  
  
    function storeString(string memory _data) public {  
        stringCount++;  
        strings[stringCount] = _data;  
        emit StringStored(stringCount, _data);  
    }  
  
    function retrieveString(uint256 _id) public view returns (string memory) {  
        require(_id > 0 && _id <= stringCount, "ID out of bounds");  
        return strings[_id];  
    }  
  
    function getStringCount() public view returns (uint256) {  
        return stringCount;  
    }  
}
