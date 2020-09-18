package idencoder

import (
    "errors"
    "fmt"
    "strconv"
)

const BASE = 36

var (
    primeMap = map[int]int64{
        8:  2199023255521,
        9:  70368744177587,
        10: 2251799813685083,
    }

    lenMap = map[int]int{
        8:  41,
        9:  46,
        10: 51,
    }
)

// Encode 将整型唯一ID转换成字符串ID
func Encode(uniqueId int64, password string, idLen int) (string, error) {
    if !checkLen(idLen) {
        return "", errors.New("idLen parameter invalid")
    }

    totalBits := lenMap[idLen]
    prime := primeMap[idLen]

    passwd := int64(encodePassword(password, totalBits))
    ret := inverse(uniqueId, prime)
    ret ^= passwd

    format := fmt.Sprintf("%%0%ds", idLen)
    return fmt.Sprintf(format, strconv.FormatInt(ret, BASE)), nil
}

// Decode 从字符串ID分解出整型ID以及地域信息
func Decode(id string, password string) (int64, error) {
    idLen := len(id)
    if !checkLen(idLen) {
        return -1, errors.New("ID length is invalid")
    }

    totalBits := lenMap[idLen]
    prime := primeMap[idLen]

    passwd := encodePassword(password, totalBits)
    v, err := strconv.ParseInt(id, BASE, 64)
    if err != nil {
        return -1, errors.New("decode id failed: " + err.Error())
    }
    vv := v ^ int64(passwd)
    return inverse(vv, prime), nil
}

func checkLen(len int) bool {
    if _, ok := primeMap[len]; ok {
        if _, ok := lenMap[len]; ok {
            return true
        }
    }
    return false
}

// encodePassword 将字符串按字符转成bit后求总和
func encodePassword(str string, retLen int) int {
    ret := 0
    sLen := len(str)
    strBytes := []byte(str)
    for i := 0; i < sLen; i++ {
        ret |= int(strBytes[sLen-i-1]) << (8 * i)
    }

    if retLen > 0 {
        ret &= ((1 << retLen) - 1)
    }
    return ret
}

func decodePassword(v int) string {
    ret := []byte{}
    for v > 0 {
        ret = append(ret, byte(v&255))
        v >>= 8
    }
    return string(reverseBytes(ret))
}

func reverseBytes(b []byte) []byte {
    for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
        b[i], b[j] = b[j], b[i]
    }
    return b
}

/**
 * 参考算法导论 推论31.26
 * 当gcd(a, b) = 1, 则方程ax + ny = 1有唯一解,这个唯一解就是 extendedEuclid(a, b) 算法得到的x
 */
func inverse(a, b int64) int64 {
    tmp := extendedEuclid(a, b)
    if tmp[1] < 0 {
        tmp[1] = tmp[1]%b + b
    }
    return tmp[1]
}

/**
 * 参考算法导论 欧几里德算法的推广形式
 * d = ax + by
 * 根据欧几里德算法推广形式计算出d,x,y
 */
func extendedEuclid(a, b int64) [3]int64 {
    if b <= 0 {
        return [3]int64{a, 1, 0}
    }
    tmp := extendedEuclid(b, a%b)
    return [3]int64{tmp[0], tmp[2], tmp[1] - ((a / b) * tmp[2])}
}
