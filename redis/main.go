package main

import (
	"context"
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"
	"github.com/spf13/pflag"
)

var rdb *redis.Client

// 命令行选项定义
var (
	addr        = pflag.StringP("addr", "", "127.0.0.1:6379", "Address of your Redis server(ip:port).")
	username    = pflag.StringP("username", "u", "", "Username for access to redis service.")
	password    = pflag.StringP("password", "p", "", "Optional auth password for redis db.")
	database    = pflag.IntP("database", "D", 0, "Database to be selected after connecting to the server.")
	maxRetries  = pflag.IntP("max-retries", "", 3, "Maximum number of retries before giving up.")
	dialTimeout = pflag.DurationP("dial-timeout", "", 5*time.Second, "Dial timeout for establishing new connections.")
	help        = pflag.BoolP("help", "h", false, "Print this help message")
)

func main() {
	// 解析命令行参数
	pflag.CommandLine.SortFlags = false
	pflag.Usage = func() {
		pflag.PrintDefaults()
	}
	pflag.Parse()
	if *help {
		pflag.Usage()
		return
	}

	// 创建一个上下文
	ctx := context.Background()

	// 创建 Redis 客户端
	rdb = redis.NewClient(&redis.Options{
		Addr:        *addr,
		Username:    *username,
		Password:    *password,
		DB:          *database,
		MaxRetries:  *maxRetries,
		DialTimeout: *dialTimeout,
	})
	// 检查客户端是否成功创建
	if _, err := rdb.Ping(ctx).Result(); err != nil {
		panic(err)
	}

	// -------------------- String 类型 ----------------------

	// 【Set】设置key value，第三个参数 0 表示不过期
	string_res, err := rdb.Set(ctx, "key1", "value1", 0).Result()
	if err != nil {
		panic(err)
	}
	fmt.Println(string_res) // OK

	// 【Set】设置一分钟过期
	string_res, err = rdb.Set(ctx, "key2", "value2", time.Minute).Result()
	if err != nil {
		panic(err)
	}
	fmt.Println(string_res) // OK

	//  下文为了精简代码，忽略错误处理部分，实际项目代码需要加上
	// 【Get】获取key对应的value
	string_res, _ = rdb.Get(ctx, "key1").Result()
	fmt.Println(string_res) // value1

	// 【GetSet】设置一个key值，并返回这个key的旧值
	string_res, _ = rdb.GetSet(ctx, "key1", "value1_new").Result()
	fmt.Println(string_res) // value1

	// 【SetNX】如果key不存在，则设置其值，不存在返回true， 已存在返回false
	bool_res, _ := rdb.SetNX(ctx, "key3", "value3", 0).Result()
	fmt.Println(bool_res) // true or false

	// 【MGet】批量查询key的值，返回数组
	array_res, _ := rdb.MGet(ctx, "key1", "key2", "key3").Result()
	fmt.Println(array_res) // [value1_new value2 value3]

	// 【MSet】批量设置key4、key5的值，常规方式
	string_res, _ = rdb.MSet(ctx, "key4", "value4", "key5", "value5").Result()
	fmt.Println(string_res) // OK

	// 【MSet】批量设置key4、key5的值，数组方式
	string_res, _ = rdb.MSet(ctx, []string{"key4", "value4", "key5", "value5"}).Result()
	fmt.Println(string_res) // OK

	// 【MSet】批量设置key4、key5的值，map方式
	string_res, _ = rdb.MSet(ctx, map[string]any{"key4": "value4", "key5": "value5"}).Result()
	fmt.Println(string_res) // OK

	// 【Incr】数值自增，key必须是数值或者不存在
	int64_res, _ := rdb.Incr(ctx, "key6").Result()
	fmt.Println(int64_res) // 1

	// 【IncrBy】数值自增，自定义增步长，key必须是数值或者不存在，每次减2
	int64_res, _ = rdb.IncrBy(ctx, "key7", -2).Result()
	fmt.Println(int64_res) // -2

	// 【IncrByFloat】数值自增，自定义增步长(浮点数)，key必须是数值或者不存在
	float64_res, _ := rdb.IncrByFloat(ctx, "key8", 6.23).Result()
	fmt.Println(float64_res) // 6.23

	// 【Decr】数值自减，key必须是数值或者不存在
	int64_res, _ = rdb.Decr(ctx, "key9").Result()
	fmt.Println(int64_res) // -1

	// 【DecrBy】数值自减，自定义步长，key必须是数值或者不存在
	int64_res, _ = rdb.DecrBy(ctx, "key10", 100).Result()
	fmt.Println(int64_res) // -100

	// 【Del】单个删除数
	int64_res, _ = rdb.Del(ctx, "key10").Result()
	fmt.Println(int64_res) // 1

	// 【Del】多个key批量删除，返回删除key 的个数
	int64_res, _ = rdb.Del(ctx, "key9", "key8", "key7", "key6").Result()
	fmt.Println(int64_res) // 4

	// 【Expire】设置key的过期时间，key不存在返回false，设置10秒过期
	bool_res, _ = rdb.Expire(ctx, "key1", time.Second*10).Result()
	fmt.Println(bool_res) // true

	// 【Exists】返回key存在的个数
	int64_res, _ = rdb.Exists(ctx, "key2", "key3", "key55").Result()
	fmt.Println(int64_res) // 2

	// -------------------- Hash 类型 ----------------------

	// 【HSet】为hash key设置字段名和值，返回不存在的field数
	int64_res, _ = rdb.HSet(ctx, "hkey1", "name", "zhangsan").Result()
	fmt.Println(int64_res) // 1

	// 【HSet】批量设置，数组模式
	int64_res, _ = rdb.HSet(ctx, "hkey1", []string{"sex", "man", "age", "18"}).Result()
	fmt.Println(int64_res) // 2

	// 【HSet】批量设置， map模式
	int64_res, _ = rdb.HSet(ctx, "hkey1", map[string]any{"name": "wangwu", "age": 100}).Result()
	fmt.Println(int64_res) // 0

	// 【HGet】获取 key 指定字段的值
	string_res, _ = rdb.HGet(ctx, "hkey1", "name").Result()
	fmt.Println(string_res) // wangwu

	// 【HGetAll】获取 key 的所有字段和值，返回 map[string]string
	map_res, _ := rdb.HGetAll(ctx, "hkey1").Result()
	fmt.Println(map_res) // map[age:100 name:wangwu sex:man]

	// 【HIncrBy】使 key 指定的字段自增 某个整数值
	int64_res, _ = rdb.HIncrBy(ctx, "hkey1", "count", 1).Result()
	fmt.Println(int64_res) // 1

	// 【HIncrByFloat】使 key 指定的字段自增 某个浮点数值
	float64_res, _ = rdb.HIncrByFloat(ctx, "hkey1", "count", 2.2).Result()
	fmt.Println(float64_res) // 2.2

	// 【HKeys】返回 key 的所有字段名，返回数组
	array_string_res, _ := rdb.HKeys(ctx, "hkey1").Result()
	fmt.Println(array_string_res) // [age name sex score count]

	// 【HLen】返回 key 的所有字段的总数
	int64_res, _ = rdb.HLen(ctx, "hkey1").Result()
	fmt.Println(int64_res) // 5

	// 【HMGet】获取 key 指定字段(多个)的值
	array_res, _ = rdb.HMGet(ctx, "hkey1", "name", "age", "score").Result()
	fmt.Println(array_res) // [wangwu 100 1]

	// 【HMSet】批量设置，map模式
	bool_res, _ = rdb.HMSet(ctx, "hkey1", map[string]any{"name": "zhaoliu", "age": 33}).Result()
	fmt.Println(bool_res) // true

	// 【HSetNX】如果字段不存在，则设置值
	bool_res, _ = rdb.HSetNX(ctx, "hkey1", "name", "da da da").Result()
	fmt.Println(bool_res) // false

	// 【HDel】删除 key 的指定字段（支持批量），返回删除字段的个数
	int64_res, _ = rdb.HDel(ctx, "hkey1", "name", "age").Result()
	fmt.Println(int64_res) // 2

	// 【HExists】检查 key 的指定字段是否存在
	bool_res, _ = rdb.HExists(ctx, "hkey1", "score").Result()
	fmt.Println(bool_res) // true

	// -------------------- List 类型 ----------------------

	// 【LPush】从左边插入数据，支持一次插入任意个数据，返回数据的总数
	int64_res, _ = rdb.LPush(ctx, "lkey1", "one", "two", "three", "four").Result()
	fmt.Println(int64_res) // 4

	// 【LPushX】用法与上述一样，但只有列表存在时才插入
	int64_res, _ = rdb.LPushX(ctx, "lkey2", "one", "two").Result()
	fmt.Println(int64_res) // 0

	// 【RPop】从列表的右边删除数据，并返回删除的数据
	string_res, _ = rdb.RPop(ctx, "lkey1").Result()
	fmt.Println(string_res) // one

	// 【RPush】、【RPushX】、【LPop】与上述对应，反向相反，不再示例

	// 【LLen】返回列表元素的总数
	int64_res, _ = rdb.LLen(ctx, "lkey1").Result()
	fmt.Println(int64_res) // 3

	// 【LRange】返回列表某范围的元素，开始索引为0，-1表示到末尾
	array_string_res, _ = rdb.LRange(ctx, "lkey1", 1, -1).Result()
	fmt.Println(array_string_res) // [three two]

	// 【LRem】从列表左边开始删除指定元素, 返回结果为成功删除的个数
	//  1表示删除从左边开始遇到的1个 two;
	//  如果是-2表示删除从右边开始遇到的两个two
	//  如果为 0 表示删除列表中的所有 two
	int64_res, _ = rdb.LRem(ctx, "lkey1", 1, "two").Result()
	fmt.Println(int64_res) // 1

	// 再插入一些测试数据
	rdb.LPush(ctx, "lkey1", "five", "six", "seven", "eight").Result()

	// 【LIndex】返回指定索引位置的元素（左侧索引从0开始）
	string_res, _ = rdb.LIndex(ctx, "lkey1", 3).Result()
	fmt.Println(string_res) // five

	// 【LInsert】在指定位置插入元素, before为指定元素之前， after为之后
	//  返回插入之后元素的数量，指定元素不存在返回 -1
	int64_res, _ = rdb.LInsert(ctx, "lkey1", "before", "six", "before_one").Result()
	fmt.Println(int64_res) // 7

	// -------------------- Set 类型 ----------------------
	//  与List 类型的区别是集合中的元素不能重复，已存在则不添加

	// 【SAdd】添加元素，支持批量，返回添加成功的个数
	int64_res, _ = rdb.SAdd(ctx, "skey1", 100, 100, 200, 300, 400, 500, 600).Result()
	fmt.Println(int64_res) // 6  两个100只算1个

	// 【SCard】获取集合元素个数
	int64_res, _ = rdb.SCard(ctx, "skey1").Result()
	fmt.Println(int64_res) // 6

	// 【SIsMember】判断元素是否在集合中
	bool_res, _ = rdb.SIsMember(ctx, "skey1", 500).Result()
	fmt.Println(bool_res) // true

	// 【SMembers】获取集合中的所有元素
	array_string_res, _ = rdb.SMembers(ctx, "skey1").Result()
	fmt.Println(array_string_res) // [100 200 300 400 500 600]

	// 【SRem】删除集合中的元素，返回被删元素的个数
	int64_res, _ = rdb.SRem(ctx, "skey1", 100, 200).Result()
	fmt.Println(int64_res) // 2

	// 【SPop】随机删除1个元素，并返回删除的元素
	string_res, _ = rdb.SPop(ctx, "skey1").Result()
	fmt.Println(string_res) // 500

	// 【SPop】随机删除n个元素，并返回删除的元素数组
	array_string_res, _ = rdb.SPopN(ctx, "skey1", 2).Result()
	fmt.Println(array_string_res) // [400 300]

	// -------------------- Sorted Set 类型 ----------------------
	//  带排序的集合，Score 为排序的权重值
	// 【ZAdd】添加元素，支持批量，返回成功添加的个数，元素存在则更新分数
	int64_res, _ = rdb.ZAdd(ctx, "zkey1",
		redis.Z{Score: 9.2, Member: "one"},
		redis.Z{Score: 8.3, Member: "two"},
		redis.Z{Score: 7.3, Member: "three"},
	).Result()
	fmt.Println(int64_res) // 3

	// 【ZCount】返回指定分数区间的元素个数（注意要用字符串），返回 7 <= 分数 <= 9的元素个数
	//  后两个参数如果是 "(7" , "(9" 则表示 7 < 分数 < 9
	int64_res, _ = rdb.ZCount(ctx, "zkey1", "7", "9").Result()
	fmt.Println(int64_res) // 2

	// 【ZIncrBy】给指定元素增加分数
	float64_res, _ = rdb.ZIncrBy(ctx, "zkey1", 2.2, "one").Result()
	fmt.Println(float64_res) // 11.399999999999999

	// 【ZRange】按分数从小到大返回元素集合，索引范围， -1表示全部
	array_string_res, _ = rdb.ZRange(ctx, "zkey1", 0, -1).Result()
	fmt.Println(array_string_res) // [three two one]

	// 【ZRevRange】按分数大到小返回元素集合，索引范围， -1表示全部
	array_string_res, _ = rdb.ZRevRange(ctx, "zkey1", 0, 1).Result()
	fmt.Println(array_string_res) // [one two]

	// 【ZRangeByScore】按分数从小到大返回元素集合， 索引范围，分数范围
	op := redis.ZRangeBy{
		Min:    "(5", // > 5
		Max:    "9",  // <= 9
		Offset: 0,    // 类似mysql的  偏移位置
		Count:  5,    // 类似mysql的 limit 数量
	}
	array_string_res, _ = rdb.ZRangeByScore(ctx, "zkey1", &op).Result()
	fmt.Println(array_string_res) // [three two]

	// 【ZRevRangeByScore】按分数从大到小返回元素集合， 索引范围，分数范围
	array_string_res, _ = rdb.ZRevRangeByScore(ctx, "zkey1", &op).Result()
	fmt.Println(array_string_res) // [two three]

	// 【ZRangeByScoreWithScores】同时返回元素和分数，从小到大
	array_struct_res, _ := rdb.ZRangeByScoreWithScores(ctx, "zkey1", &op).Result()
	fmt.Println(array_struct_res) // [{7.3 three} {8.3 two}]

	// 【ZRevRangeByScoreWithScores】同时返回元素和分数，从大到小
	array_struct_res, _ = rdb.ZRevRangeByScoreWithScores(ctx, "zkey1", &op).Result()
	fmt.Println(array_struct_res) // [{8.3 two} {7.3 three}]

	// 【ZRem】删除集合元素，支持批量，返回成功删除的个数
	int64_res, _ = rdb.ZRem(ctx, "zkey1", "one", "two").Result()
	fmt.Println(int64_res) // 2

	// 【ZRemRangeByRank】按索引范围删除，返回成功删除的个数
	int64_res, _ = rdb.ZRemRangeByRank(ctx, "zkey1", 0, 1).Result()
	fmt.Println(int64_res) // 1

	// 【ZRemRangeByScore】按分数范围删除，返回成功删除的个数
	int64_res, _ = rdb.ZRemRangeByScore(ctx, "zkey1", "(2", "8").Result()
	fmt.Println(int64_res) // 0

	rdb.ZAdd(ctx, "zkey1",
		redis.Z{Score: 1.1, Member: "one"},
		redis.Z{Score: 2.2, Member: "two"},
		redis.Z{Score: 3.3, Member: "three"},
	).Result()

	// 【ZScore】查询指定元素的分数，不存在返回0
	float64_res, _ = rdb.ZScore(ctx, "zkey1", "one").Result()
	fmt.Println(float64_res) // 1.1

	// 【ZRank】查询指定元素从小到大的排名，索引从0开始
	int64_res, _ = rdb.ZRank(ctx, "zkey1", "two").Result()
	fmt.Println(int64_res) // 1

	// 【ZRevRank】查询指定元素从大到小的排名，索引从0开始
	int64_res, _ = rdb.ZRevRank(ctx, "zkey1", "one").Result()
	fmt.Println(int64_res) // 2

	// -------------------- Bitmap 类型 ----------------------
	// 【SetBit】设置位
	int64_res, _ = rdb.SetBit(ctx, "mybitmap", 0, 1).Result()
	fmt.Println(int64_res) // 0

	// 【GetBit】获取位
	int64_res, _ = rdb.GetBit(ctx, "mybitmap", 0).Result()
	fmt.Println(int64_res) // 1

	// 【BitCount】统计位数量
	int64_res, _ = rdb.BitCount(ctx, "mybitmap", nil).Result()
	fmt.Println(int64_res) // 1

	// -------------------- HyperLogLog 类型 ----------------------
	// 【PFAdd】添加元素到 HyperLogLog
	int64_res, _ = rdb.PFAdd(ctx, "hll_key", "item1", "item2", "item3").Result()
	fmt.Println(int64_res) // 1

	// 【PFCount】获取 HyperLogLog 的基数（不重复元素的数量）
	int64_res, _ = rdb.PFCount(ctx, "hll_key").Result()
	fmt.Println(int64_res) // 3

	// -------------------- Geospatial 类型 ----------------------
	// 【GeoAdd】添加地理空间位置
	int64_res, _ = rdb.GeoAdd(ctx, "locations", &redis.GeoLocation{Longitude: 13.361389, Latitude: 38.115556, Name: "Palermo"}).Result()
	fmt.Println(int64_res) // 1

	// 【GeoDist】获取两个地点之间的距离
	float64_res, _ = rdb.GeoDist(ctx, "locations", "Palermo", "Catania", "km").Result()
	fmt.Println(float64_res) // 0

	// -------------------- 事务处理 ----------------------
	// 类似于mysql事务，事务中的所有命令会按顺序执行，并且不被外界命令所打断
	// 原子性：要么全部执行，要么全部不执行
	// 【TxPipeline】开启一个事务
	pipe := rdb.TxPipeline()
	p := pipe.Set(ctx, "pkey1", 1, 0)
	pipe.IncrBy(ctx, "pkey1", 2)
	// 【Exec】执行事务，这一步之后才开始顺序执行所有命令
	res, err := pipe.Exec(ctx)
	fmt.Println(res, err) // [set pkey1 1: OK incrby pkey1 2: 3] <nil>
	fmt.Println(p.Val())  // OK

	// -------------------- 消息订阅 ----------------------
	Subscribe()

	select {}
}

// 订阅
func Subscribe() {
	// 模拟发布者，每两秒往频道里发送1次消息
	go func() {
		ctx := context.Background()
		c := time.Tick(time.Second * 2)
		var count int
		for {
			<-c
			count++
			// 【Publish】往指定频道发送消息
			rdb.Publish(ctx, "channel1", fmt.Sprintf("hello: %d", count))
		}
	}()

	var receive func(int) = func(num int) {
		ctx := context.Background()
		//【Subscribe】订阅频道
		sub := rdb.Subscribe(ctx, "channel1")
		// 订阅者实时接收频道中的消息
		for msg := range sub.Channel() {
			// 打印频道号和消息内容
			fmt.Printf("我是接收者%d，我接收到来自频道%s的消息: %s\n",
				num, msg.Channel, msg.Payload)
		}

	}

	// 模拟接收者1
	go receive(1)
	// 模拟接收者2
	go receive(2)

}
