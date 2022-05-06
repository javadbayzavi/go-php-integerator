/*
*** @author: Javad Bayzavi
*** @version: 1.0.1
*** @email: javadbayzavi@gmail.com
*** @year: 2021
*/

package Factory

import (
	"go-microservices/lib/Core/Source"
	"go-microservices/lib/Providers/Db/MySql"
	"go-microservices/lib/Providers/Interface"
)

func CreateProvider() Interface.ProviderInterface {

	switch Source.SOURCE_CURRENT_PROVIDER {
	case "my-sql":
		//var temp = createMySqlProvider()
		return createMySqlProvider()

	}

	//Default data provider is MySql
	return createMySqlProvider()

}

func createMySqlProvider() *MySql.MySqlProvider {
	//Initialize provider
	provider := &MySql.MySqlProvider{}

	//Set provider connection settings
	provider.SetDatasource(Source.SOURCE_USER)
	provider.SetHost(Source.SOURCE_HOST)
	provider.SetPassword(Source.SOURCE_PASSWORD)
	provider.SetUser(Source.SOURCE_USER)

	return provider
}
