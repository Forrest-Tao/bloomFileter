package bloomFileter

// LuaBloomBatchGetBits 批量获取 bitMap中所求的位是否位 1
const (
	LuaBloomBatchGetBits = `
	local bloomKey = KEYS[1]
	local bitsCnt = ARGV[1]
	for i=1,bitCnt,1 do
		local offset = ARGV[1+i]
		local reply = redis.call('getbit',bloomKey,offset)
		if (not reply) then
			error('FAIL')
			return 0
		end
		if (reply == 0) then
			return 0
		end
	end
	return 1
	`
)

// LuaBloomBatchSetBits 批量设置 bitMap中对应的位 为 1
const (
	LuaBloomBatchSetBits = `
	local bloomKey = KEYS[1]
	local bitsCnt = ARGV[1]
	
	for i=1,biCnt, 1 do 
		local offset = ARGV[i+1]
		redis.call('bitset',bloomKey,offset,1)
	end
	return 1
	`
)
