package ncloud

import (
	"fmt"
	"regexp"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/acctest"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
)

func TestAccResourceNcloudBlockStorage_classic_basic(t *testing.T) {
	var storageInstance BlockStorage
	name := fmt.Sprintf("tf-storage-basic-%s", acctest.RandString(5))
	resourceName := "ncloud_block_storage.storage"

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccClassicProviders,
		CheckDestroy: func(state *terraform.State) error {
			return testAccCheckBlockStorageDestroyWithProvider(state, testAccClassicProvider)
		},
		Steps: []resource.TestStep{
			{
				Config: testAccBlockStorageClassicConfig(name),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckBlockStorageExistsWithProvider(resourceName, &storageInstance, testAccClassicProvider),
					resource.TestMatchResourceAttr(resourceName, "id", regexp.MustCompile(`^\d+$`)),
					resource.TestCheckResourceAttr(resourceName, "name", name),
					resource.TestCheckResourceAttr(resourceName, "status", "ATTAC"),
					resource.TestCheckResourceAttr(resourceName, "size", "10"),
					resource.TestCheckResourceAttr(resourceName, "type", "SVRBS"),
					resource.TestCheckResourceAttr(resourceName, "disk_type", "NET"),
					resource.TestMatchResourceAttr(resourceName, "block_storage_no", regexp.MustCompile(`^\d+$`)),
					resource.TestMatchResourceAttr(resourceName, "server_instance_no", regexp.MustCompile(`^\d+$`)),
					resource.TestCheckResourceAttr(resourceName, "description", ""),
					resource.TestCheckResourceAttr(resourceName, "device_name", "/dev/xvdb"),
					resource.TestCheckResourceAttr(resourceName, "product_code", "SPBSTBSTAD000002"),
					resource.TestCheckResourceAttr(resourceName, "disk_detail_type", "HDD"),
				),
			},
			{
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccResourceNcloudBlockStorage_vpc_basic(t *testing.T) {
	var storageInstance BlockStorage
	name := fmt.Sprintf("tf-storage-basic-%s", acctest.RandString(5))
	resourceName := "ncloud_block_storage.storage"

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		CheckDestroy: func(state *terraform.State) error {
			return testAccCheckBlockStorageDestroyWithProvider(state, testAccProvider)
		},
		Steps: []resource.TestStep{
			{
				Config: testAccBlockStorageVpcConfig(name),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckBlockStorageExistsWithProvider(resourceName, &storageInstance, testAccProvider),
					resource.TestMatchResourceAttr(resourceName, "id", regexp.MustCompile(`^\d+$`)),
					resource.TestCheckResourceAttr(resourceName, "name", name),
					resource.TestCheckResourceAttr(resourceName, "status", "ATTAC"),
					resource.TestCheckResourceAttr(resourceName, "size", "10"),
					resource.TestCheckResourceAttr(resourceName, "type", "SVRBS"),
					resource.TestCheckResourceAttr(resourceName, "disk_type", "NET"),
					resource.TestMatchResourceAttr(resourceName, "block_storage_no", regexp.MustCompile(`^\d+$`)),
					resource.TestMatchResourceAttr(resourceName, "server_instance_no", regexp.MustCompile(`^\d+$`)),
					resource.TestCheckResourceAttr(resourceName, "description", ""),
					resource.TestCheckResourceAttr(resourceName, "device_name", "/dev/xvdb"),
					resource.TestCheckResourceAttr(resourceName, "product_code", "SPBSTBSTAD000006"),
					resource.TestCheckResourceAttr(resourceName, "disk_detail_type", "SSD"),
				),
			},
			{
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccResourceNcloudBlockStorage_classic_ChangeServerInstance(t *testing.T) {
	var storageInstance BlockStorage
	name := fmt.Sprintf("tf-storage-update-%s", acctest.RandString(5))
	resourceName := "ncloud_block_storage.storage"

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccClassicProviders,
		CheckDestroy: func(state *terraform.State) error {
			return testAccCheckBlockStorageDestroyWithProvider(state, testAccClassicProvider)
		},
		Steps: []resource.TestStep{
			{
				Config: testAccBlockStorageClassicConfigUpdate(name, "ncloud_server.foo.id"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckBlockStorageExistsWithProvider(resourceName, &storageInstance, testAccClassicProvider),
				),
			},
			{
				Config: testAccBlockStorageClassicConfigUpdate(name, "ncloud_server.bar.id"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckBlockStorageExistsWithProvider(resourceName, &storageInstance, testAccClassicProvider),
				),
			},
		},
	})
}

func TestAccResourceNcloudBlockStorage_vpc_ChangeServerInstance(t *testing.T) {
	var storageInstance BlockStorage
	name := fmt.Sprintf("tf-storage-update-%s", acctest.RandString(5))
	resourceName := "ncloud_block_storage.storage"

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		CheckDestroy: func(state *terraform.State) error {
			return testAccCheckBlockStorageDestroyWithProvider(state, testAccProvider)
		},
		Steps: []resource.TestStep{
			{
				Config: testAccBlockStorageVpcConfigUpdate(name, "ncloud_server.foo.id"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckBlockStorageExistsWithProvider(resourceName, &storageInstance, testAccProvider),
				),
			},
			{
				Config: testAccBlockStorageVpcConfigUpdate(name, "ncloud_server.bar.id"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckBlockStorageExistsWithProvider(resourceName, &storageInstance, testAccProvider),
				),
			},
		},
	})
}

func TestAccResourceNcloudBlockStorage_classic_size(t *testing.T) {
	var storageInstance BlockStorage
	name := fmt.Sprintf("tf-storage-size-%s", acctest.RandString(5))
	resourceName := "ncloud_block_storage.storage"

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccClassicProviders,
		CheckDestroy: func(state *terraform.State) error {
			return testAccCheckBlockStorageDestroyWithProvider(state, testAccClassicProvider)
		},
		Steps: []resource.TestStep{
			{
				Config: testAccBlockStorageClassicConfigWithSize(name, 10),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckBlockStorageExistsWithProvider(resourceName, &storageInstance, testAccClassicProvider),
					resource.TestCheckResourceAttr(resourceName, "size", "10"),
				),
			},
			{
				Config: testAccBlockStorageClassicConfigWithSize(name, 20),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckBlockStorageExistsWithProvider(resourceName, &storageInstance, testAccClassicProvider),
					resource.TestCheckResourceAttr(resourceName, "size", "20"),
				),
			},
			{
				Config: testAccBlockStorageClassicConfigWithSize(name, 10),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckBlockStorageExistsWithProvider(resourceName, &storageInstance, testAccClassicProvider),
				),
				ExpectError: regexp.MustCompile("The storage size is only expandable, not shrinking. new size(\\d+) must be greater than the existing size(\\d+)"),
			},
		},
	})
}

func TestAccResourceNcloudBlockStorage_vpc_size(t *testing.T) {
	var storageInstance BlockStorage
	name := fmt.Sprintf("tf-storage-size-%s", acctest.RandString(5))
	resourceName := "ncloud_block_storage.storage"

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		CheckDestroy: func(state *terraform.State) error {
			return testAccCheckBlockStorageDestroyWithProvider(state, testAccProvider)
		},
		Steps: []resource.TestStep{
			{
				Config: testAccBlockStorageVpcConfigWithSize(name, 10),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckBlockStorageExistsWithProvider(resourceName, &storageInstance, testAccProvider),
					resource.TestCheckResourceAttr(resourceName, "size", "10"),
				),
			},
			{
				Config: testAccBlockStorageVpcConfigWithSize(name, 20),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckBlockStorageExistsWithProvider(resourceName, &storageInstance, testAccProvider),
					resource.TestCheckResourceAttr(resourceName, "size", "20"),
				),
			},
			{
				Config: testAccBlockStorageVpcConfigWithSize(name, 10),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckBlockStorageExistsWithProvider(resourceName, &storageInstance, testAccClassicProvider),
				),
				ExpectError: regexp.MustCompile("new size(\\d+) must be greater than the existing size(\\d+)"),
			},
		},
	})
}

func testAccCheckBlockStorageExistsWithProvider(n string, i *BlockStorage, provider *schema.Provider) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("no ID is set")
		}

		config := provider.Meta().(*ProviderConfig)
		storage, err := getBlockStorage(config, rs.Primary.ID)
		if err != nil {
			return nil
		}

		if storage != nil {
			*i = *storage
			return nil
		}

		return fmt.Errorf("block storage not found")
	}
}

func testAccCheckBlockStorageDestroyWithProvider(s *terraform.State, provider *schema.Provider) error {
	config := provider.Meta().(*ProviderConfig)

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "ncloud_block_storage" {
			continue
		}
		blockStorage, err := getBlockStorage(config, rs.Primary.ID)

		if blockStorage == nil {
			continue
		}
		if err != nil {
			return err
		}
		if blockStorage != nil && *blockStorage.Status != "ATTAC" {
			return fmt.Errorf("found attached block storage: %s", *blockStorage.BlockStorageInstanceNo)
		}
	}

	return nil
}

func testAccBlockStorageClassicConfigWithSize(name string, size int) string {
	return fmt.Sprintf(`
resource "ncloud_login_key" "loginkey" {
	key_name = "%[1]s-key"
}

resource "ncloud_server" "server" {
	name = "%[1]s"
	server_image_product_code = "SPSW0LINUX000032"
	server_product_code = "SPSVRSTAND000004"
	login_key_name = "${ncloud_login_key.loginkey.key_name}"
}

resource "ncloud_block_storage" "storage" {
	server_instance_no = ncloud_server.server.id
	name = "%[1]s"
	size = "%[2]d"
}
`, name, size)
}

func testAccBlockStorageClassicConfig(name string) string {
	return testAccBlockStorageClassicConfigWithSize(name, 10)
}

func testAccBlockStorageVpcConfigWithSize(name string, size int) string {
	return fmt.Sprintf(`
resource "ncloud_login_key" "loginkey" {
	key_name = "%[1]s-key"
}

resource "ncloud_vpc" "test" {
	name               = "%[1]s"
	ipv4_cidr_block    = "10.5.0.0/16"
}

resource "ncloud_subnet" "test" {
	vpc_no             = ncloud_vpc.test.vpc_no
	name               = "%[1]s"
	subnet             = "10.5.0.0/24"
	zone               = "KR-2"
	network_acl_no     = ncloud_vpc.test.default_network_acl_no
	subnet_type        = "PUBLIC"
	usage_type         = "GEN"
}

resource "ncloud_server" "server" {
	subnet_no = ncloud_subnet.test.id
	name = "%[1]s"
	server_image_product_code = "SW.VSVR.OS.LNX64.CNTOS.0703.B050"
	server_product_code = "SVR.VSVR.STAND.C002.M008.NET.HDD.B050.G002"
	login_key_name = ncloud_login_key.loginkey.key_name
}

resource "ncloud_block_storage" "storage" {
	server_instance_no = ncloud_server.server.id
	name = "%[1]s"
	size = "%[2]d"
}
`, name, size)
}

func testAccBlockStorageVpcConfig(name string) string {
	return testAccBlockStorageVpcConfigWithSize(name, 10)
}
func testAccBlockStorageClassicConfigUpdate(name, serverInstanceNo string) string {
	return fmt.Sprintf(`
resource "ncloud_login_key" "loginkey" {
	key_name = "%[1]s-key"
}

resource "ncloud_server" "foo" {
	name = "%[1]s-foo"
	server_image_product_code = "SPSW0LINUX000032"
	server_product_code = "SPSVRSTAND000004"
	login_key_name = "${ncloud_login_key.loginkey.key_name}"
}

resource "ncloud_server" "bar" {
	name = "%[1]s-bar"
	server_image_product_code = "SPSW0LINUX000032"
	server_product_code = "SPSVRSTAND000004"
	login_key_name = "${ncloud_login_key.loginkey.key_name}"
}

resource "ncloud_block_storage" "storage" {
	server_instance_no =  %[2]s
	name = "%[1]s"
	size = "10"
}
`, name, serverInstanceNo)
}

func testAccBlockStorageVpcConfigUpdate(name, serverInstanceNo string) string {
	return fmt.Sprintf(`
resource "ncloud_login_key" "loginkey" {
	key_name = "%[1]s-key"
}

resource "ncloud_vpc" "test" {
	name               = "%[1]s"
	ipv4_cidr_block    = "10.5.0.0/16"
}

resource "ncloud_subnet" "test" {
	vpc_no             = ncloud_vpc.test.vpc_no
	name               = "%[1]s"
	subnet             = "10.5.0.0/24"
	zone               = "KR-2"
	network_acl_no     = ncloud_vpc.test.default_network_acl_no
	subnet_type        = "PUBLIC"
	usage_type         = "GEN"
}

resource "ncloud_server" "foo" {
	subnet_no = ncloud_subnet.test.id
	name = "%[1]s-foo"
	server_image_product_code = "SW.VSVR.OS.LNX64.CNTOS.0703.B050"
	server_product_code = "SVR.VSVR.STAND.C002.M008.NET.HDD.B050.G002"
	login_key_name = ncloud_login_key.loginkey.key_name
}

resource "ncloud_server" "bar" {
	subnet_no = ncloud_subnet.test.id
	name = "%[1]s-bar"
	server_image_product_code = "SW.VSVR.OS.LNX64.CNTOS.0703.B050"
	server_product_code = "SVR.VSVR.STAND.C002.M008.NET.HDD.B050.G002"
	login_key_name = ncloud_login_key.loginkey.key_name
}

resource "ncloud_block_storage" "storage" {
	server_instance_no = %[2]s
	name = "%[1]s"
	size = "10"
}
`, name, serverInstanceNo)
}
