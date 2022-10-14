package cmdprefix

const (
	chmodRecursiveFormat    = "chmod -R %s"    // "chmod -R %s" chmod -R 777 /dir1
	chmodNonRecursiveFormat = "chmod %s"       // "chmod %s"chmod -R 777 /dir1
	chownRecursiveFormat    = "chown -R %s:%s" // "chown -R %s:%s" chown -R pdns:root /dir1 , chown -R username:group directory
	chownNonRecursiveFormat = "chown %s:%s"    // "chown %s:%s" chown pdns:root /dir1 , chown username:group directory
	// https://linuxize.com/post/chgrp-command-in-linux/
	changeGroupNonRecursiveFormat = "chgrp %s"    // "chgrp %s" chgrp $group /dir1
	changeGroupRecursiveFormat    = "chgrp -R %s" // "chgrp -R %s" chgrp -R $group /dir1
	Touch                         = "touch"
	HyphenE                       = "-e"
	HyphenA                       = "-a"
	HyphenG                       = "-G"
	Chown                         = "chown"
	ChGroup                       = "chgrp"
	Root                          = "root"
	UserMod                       = "usermod"
)
