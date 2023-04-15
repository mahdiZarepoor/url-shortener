# A simple link shortener using golang + redis 




## further development  :
- َUse sha256 + base 58 encoding (more reliable algorithm to shorten the link)
- Use mysql to  store cold data and redis for hot data . (move expired data into mysql)
- Add unit tests, integration tests.
- Refactor repository layer, add service layer and use interfaces.
- Change it in a way that generates different links for different user
