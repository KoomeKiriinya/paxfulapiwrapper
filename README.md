
# Paxful API wrapper: WIP 
### NB this library/wrapper has not been published and can only list your paxful transactions

---
```markdown

To use this wrapper copy the paxful folder to your project and import it 

into your main or desired package and replace github.com/paxfulapiwrapper

with your go module. 

Also add variables in env.example to your envs.

```
---
---
```markdown

package main

import (
	"fmt"

	"github.com/paxfulapiwrapper/paxful"
)

func main() {

	pax_client := paxful.NewClient()
	transactions, err := pax_client.Transactions("1", "10", "all", "USDT")
	if err != nil {
		fmt.Println("error :" + err.Error())
	}
	fmt.Println(transactions)

}
```
---
