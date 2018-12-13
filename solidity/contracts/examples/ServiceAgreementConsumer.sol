pragma solidity ^0.4.24;

import "../Chainlinked.sol";

contract ServiceAgreementConsumer is Chainlinked {
  uint256 constant private LINK_DIVISIBILITY = 10**18;

  bytes32 internal sAId;
  bytes32 public currentPrice;

  constructor(address _link, address _coordinator, bytes32 _sAId) public {
    setLinkToken(_link);
    setOracle(_coordinator);
    sAId = _sAId;
  }

  function requestEthereumPrice(string _currency) public {
    ChainlinkLib.Run memory run = newRun(sAId, this, this.fulfill.selector);
    run.add("url", "https://min-api.cryptocompare.com/data/price?fsym=ETH&tsyms=USD,EUR,JPY");
    run.add("path", _currency);
    chainlinkRequest(run, LINK_DIVISIBILITY);
  }

  function fulfill(bytes32 _requestId, bytes32 _price)
    public
    checkChainlinkFulfillment(_requestId)
  {
    currentPrice = _price;
  }
}
