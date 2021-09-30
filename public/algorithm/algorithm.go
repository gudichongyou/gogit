package algorithm

import (
	"fmt"
	"strconv"
)

type SMPkey struct {
	Key   string
	Stype byte
}

func SortMap(mapa []map[string]interface{}, skey ...SMPkey) []map[string]interface{} {
	// rtmapa := make([]map[string]interface{}, len(mapa))
	//var trarryin []int = mak([]int, len(len(mapa)))
	lmapa := len(mapa)
	for ind := range mapa {
		for ind2 := ind + 1; ind2 < lmapa; ind2++ {
		for1:
			for indskey := range skey {
				var sskey string = skey[indskey].Key
				stype := skey[indskey].Stype
				// fmt.Println(sskey, stype)
				switch mapa[ind][sskey].(type) {
				case int64:
					if mapa[ind][sskey].(int64) > mapa[ind2][sskey].(int64) {
						if stype == 0 { //升序
							mapa[ind], mapa[ind2] = mapa[ind2], mapa[ind]
							fmt.Println("交换:", mapa[ind], mapa[ind2])
							break for1
						} else {
							//降序
							break for1

						}

					} else if mapa[ind][sskey].(int64) == mapa[ind2][sskey].(int64) {
						continue

					} else {
						if stype == 0 { //升序
							break for1

						} else { //降序
							mapa[ind], mapa[ind2] = mapa[ind2], mapa[ind]
							break for1
						}
					}
				case float64:
					if mapa[ind][sskey].(float64) > mapa[ind2][sskey].(float64) {
						if stype == 0 { //升序
							mapa[ind], mapa[ind2] = mapa[ind2], mapa[ind]
							break for1
						} else {
							//降序
							break for1

						}

					} else if mapa[ind][sskey].(float64) == mapa[ind2][sskey].(float64) {
						continue

					} else {
						if stype == 0 { //升序
							break for1

						} else { //降序
							mapa[ind], mapa[ind2] = mapa[ind2], mapa[ind]
							break for1
						}
					}
				case string:
					if mapa[ind][sskey].(string) > mapa[ind2][sskey].(string) {
						if stype == 0 { //升序
							mapa[ind], mapa[ind2] = mapa[ind2], mapa[ind]
							break for1
						} else {
							//降序
							break for1

						}

					} else if mapa[ind][sskey].(string) == mapa[ind2][sskey].(string) {
						continue

					} else {
						if stype == 0 { //升序
							break for1

						} else { //降序
							mapa[ind], mapa[ind2] = mapa[ind2], mapa[ind]
							break for1
						}
					}

				}

			}

		}

	}
	return mapa
}

func TestSMP() {
	var mpa []map[string]interface{} = make([]map[string]interface{}, 0)
	var i int64
	for i = 0; i < 10; i++ {
		func(ii int64) {
			var maps map[string]interface{} = make(map[string]interface{})
			maps["testid"] = ii
			maps["deviceid"] = "deviceid" + strconv.FormatInt(i%5, 10)
			maps["sno"] = "sno" + strconv.FormatInt(10-i, 10)
			mpa = append(mpa, maps)
		}(i)

	}
	fmt.Println(mpa)
	fmt.Println("---------")
	var skey [2]SMPkey
	skey[0] = SMPkey{Key: "deviceid", Stype: 1}
	skey[1] = SMPkey{Key: "sno", Stype: 0}
	mpa = SortMap(mpa, skey[:]...)
	fmt.Println(mpa)
}
