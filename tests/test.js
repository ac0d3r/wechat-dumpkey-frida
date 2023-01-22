var path = ObjC.classes.NSString.stringWithString_("tests/setting_db.data");
var key = ObjC.classes.PBCoder["+ decodeObjectOfClass:fromFile:"](ObjC.classes.DBEncryptInfo, path)
var data = key['- m_dbEncryptKey']();
hexdump(data.bytes(), { offset: 0, length: data.length(), header: false, ansi: false });
