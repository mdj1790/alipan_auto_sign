package config

import (
	"os"

	"gopkg.in/yaml.v3"
)

var ConfigInstance Config

var UserAgent = "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/123.0.0.0 Safari/537.36 Edg/123.0.0.0"

type Config struct {
	PushPlusToken  string `yaml:"pushplus_token"`
	RefreshToken   string `yaml:"refresh_token"`
	BilibiliCookie string `yaml:"bilibili_cookie"`
	KKCookie       string `yaml:"b-user-id=d0db50d9-439c-cc2c-e237-50ff50efe270; b-user-id=d0db50d9-439c-cc2c-e237-50ff50efe270; _UP_A4A_11_=wb9c5133aedb4f0d9263231e66dc7681; __wpkreporterwid_=f3a3d197-c812-4d81-28a6-085dcdb47c73; __itrace_wid=faa2c728-ea26-42e7-974c-0780b42e311b; _UP_335_2B_=1; __sdid=AARuGTZPKEB9Cr5moWIJ3/LnTsxPIG5sPC9mSpVxWV9gkHBMKHYHloZTfAx8LjNFa28=; ctoken=pAMJCbAuL_O7qOHPCsFOJ70o; web-grey-id=b460460e-a2b8-5dad-f1f8-c39747fedc94; web-grey-id.sig=9XpDtQqFjqpUEenb14vnm8Lrab22gikuRhmvo_V9zZw; grey-id=699a5c54-b7bb-0152-28fb-d79c49805b38; grey-id.sig=TgZYoI-8JOdhaaW_65bVH_TDe5v1qwxvumDElgCTvpU; isQuark=false; isQuark.sig=DWPHMZYiiwQ-v58AbcP-rBdSIpzO8ZnrD67BdJuPatU; _UP_30C_6A_=st9c9620114bud7eaootkvh7x5kwilkw; _UP_TS_=sg1bad0983de51c6520524aaa625dda0309; _UP_E37_B7_=sg1bad0983de51c6520524aaa625dda0309; _UP_TG_=st9c9620114bud7eaootkvh7x5kwilkw; _UP_D_=pc; __chkey=; isg=BHZ2n0MgAuf6Gfk6PAp5WECGx6x4l7rRqcJt3OBfc9n0Ixa9SCO14d1QP_9Pi7Lp; tfstk=g17KUeZEOR2ndBbps6riqxfYdlVgmlfFB95jrLvnVOBOaTKkVHAHy36BIwxkRpXt2QdmADZEAdtJNTJ5Re6uyL12Q2XuY30-TNBxr22UY_gJrUw0nr4cTXYyP-AuYu8yLQ56F5xIxSqoHUw0nPmi1HyHPTj7G3yO1dRJNBtCFha6QIpWPptSCcOXCUTWV4G61BRmOLOWO5C6QQTWFTT76F9wNbNuedGBE4I4xNP-v4AEP4_99hZGOdiWQNdpVK1CJ4g5cBKph69T3W6-1h1y2alry3CC0TRfdvwJUNC1Ri6bIDvCDQCc26aI5ebFOit5l8o5tFvp5wsxN419JMtPPhHLBK_F1Z86jrNA1NWhTNCoNzOGnd1FRUUbgeLBdFdNrRuBeGs5-H8rd8pPW1_C2g8PoZdbVxvvZDNT60oyACrB_CmqPOMA1Cp09znr4hPw6KVOk0oyAiO9n5VS40-Gm; _UP_F7E_8D_=OQNlgAklx5qlL3as0ULacEPwKLOVbxJPcg0RzQPI6KmBtV6ZMgPh38l93pgubgHDQqhaZ2Sfc0qv%2BRantbfg1mWGAUpRMP4RqXP78Wvu%2FCfvkWWGc5NhCTV71tGOIGgDBR3%2Bu6%2Fjj46CcaOUk%2FXNWE2yEx2v2eT8vOTidzNw8s%2FWtKAIxWbnCzZn4%2FJMBUub0OScUYeEhutiwvwuRNnRsR9t%2BT3RpJCeiA6u93MtuAW1kGun4123xxU4wR0Pq7NklczEGdRq2nIAcu7v22Uw2o%2FxMY0xBdeC9Korm5%2FNHnxl6K%2Bd6FXSoT9a3XIMQO359auZPiZWzrNlZe%2BqnOahXcx7KAhQIRqSOapSmL4ygJor4r5isJhRuDoXy7vJAVuH%2FRDtEJJ8rZTq0BdC23Bz%2B0MrsdgbK%2BiW; __pus=e1ec381d35b509ccfe5346691efb8081AATWBiPyqFAE2zpBNtipiRSmPifI9inDv0YVYgaAiHPQWAy5PdeTYzT10CqQaFIutuLCoiPouoGsXD4INIg+Lt1y; __kp=a2e9c750-3214-11f0-9e58-739453cabcaa; __kps=AATpVI3WEDK20qXu2r4ti5vY; __ktd=R9scYRjBRnDj3FAVa0i9iQ==; __uid=AATpVI3WEDK20qXu2r4ti5vY; __puus=fecb3804896ba1ba977b1f4e88b47c6eAARRZaDf+w0IQCEiY4IvA0WYcctWz8BVtDNsCwpQ+6vZjitjJjFNxo4UpbEPNBdvSHyxXa4yMsRZXHUfIPQxxq1rCOAVg0pp3sU72LEk2/zfO0S3aYjcy8JV7yHkTUZ86NM3uVuYn6uq1KdK1h+C4gT8p9Z++cJkKzkBIN7xWU1iuXHMYCTL32z7D2yLpYrD7ofJQ+uMydZH9JiylvKDZYoT"`
	JdCookie       string `yaml:"jd_cookie"`
}

func init() {
	LoadConfig()
}

func LoadConfig() {
	confFIle, err := os.ReadFile("./config.yaml")
	if err != nil {
		panic(err.Error())
	}
	config := Config{}
	yaml.Unmarshal(confFIle, &config)
	ConfigInstance = config
}
