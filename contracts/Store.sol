pragma solidity ^0.5.5;

contract Store {

  event dataSetter(string _nearest_pole_id, string _latitude, string _longitude, string _timestamp);
    
  string public latitude; // encrypted aadhar info
  string public longitude; // encrypted aadhar info
  string public timestamp;
  string public nearest_pole_id;

  function setData(string calldata _nearest_pole_id, string calldata _latitude, string calldata _longitude, string calldata _timestamp) external {
    latitude = _latitude;
    longitude = _longitude;
    nearest_pole_id = _nearest_pole_id;
    timestamp = _timestamp;
    emit dataSetter(latitude, longitude, nearest_pole_id, timestamp);
  }
}