var key = ObjC.chooseSync(ObjC.classes.DBEncryptInfo)[0];
var data = key['- m_dbEncryptKey']();
console.log(hexdump(data.bytes(), { offset: 0, length: data.length(), header: false, ansi: false }));
