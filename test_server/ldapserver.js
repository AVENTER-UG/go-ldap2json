var ldap = require('./ldapjs');

var server = ldap.createServer();

server.search('ou=users, o=company', function(req, res, next) {

  console.log("Search: " + req.dn.toString())

  var obj = {
    dn: req.dn.toString(),
    attributes: {
      objectclass: ['person', 'top'],
      o: 'company',
      cn: 'andreas',
      sn: 'peters',
      email: 'nope@<blub>',
    }
  };

  if (req.filter.matches(obj.attributes))
    res.send(obj);

  res.end();
});

server.listen(1389, function() {
  console.log('LDAP server listening at %s', server.url);
});
