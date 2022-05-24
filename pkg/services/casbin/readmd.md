Custom config
```go
(function() {
  return {
    /**
     * Here is custom functions for Casbin.
     * Currently, there are built-in globMatch, keyMatch, keyMatch2, keyMatch3, keyMatch4, regexMatch, ipMatch.
     */
    functions: {},
    /**
     * If the value is undefined, the Casbin does not use it.
     * example:
     * matchingForGFunction: 'globMatch'
     * matchingDomainForGFunction: 'keyMatch'
     */
    matchingForGFunction: 'keyMatch',
    matchingDomainForGFunction: 'keyMatch'
  };
})();
```

model
```ini
[request_definition]
r = sub, dom, obj, act

[policy_definition]
p = sub, obj, act
[role_definition]
g = _, _, _

[policy_effect]
e = some(where (p.eft == allow))

[matchers]
m = g(r.sub, p.sub, r.dom) && r.obj == p.obj && r.act == p.act
```

policy
```cvs
p, role_admin,  /user, read
p, role_admin, /pod, write
p, role_admin, /service, read
p, role_admin, /hello, write

g, alice, admin, /1/*
g, bob, admin, /1/2
```

Request
```cvs
role_alice, /1/1, /user, read
role_alice, /1/2, /user, read
user_bob, /1/1, /user, read
user_bob, /1/2, /user, read
```

