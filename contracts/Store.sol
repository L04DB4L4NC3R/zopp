pragma solidity ^0.5.5;

contract Store {

  event dataSetter(string _nearest_pole_id, string _latitude, string _longitude, string _timestamp, uint8 _ageSuggestive);
    
  string public latitude; // encrypted aadhar info
  string public longitude; // encrypted aadhar info
  string public timestamp;
  string public nearest_pole_id;
  uint8 public ageSuggestive;

  function setData(string calldata _nearest_pole_id, string calldata _latitude, string calldata _longitude, string calldata _timestamp, uint8 _ageSuggestive) external {
    latitude = _latitude;
    longitude = _longitude;
    nearest_pole_id = _nearest_pole_id;
    timestamp = _timestamp;
    ageSuggestive = _ageSuggestive;
  }
}