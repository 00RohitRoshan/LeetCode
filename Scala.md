```Scala
object Solution {
    def twoSum(nums: Array[Int], target: Int): Array[Int] = {
        func(nums.zipWithIndex, target)
    }
     //@scala.annotation.tailrec
    def func(numsIdx: Array[(Int, Int)], target: Int): Array[Int] = {
        val expect = target - numsIdx.head._1
        for (i <- numsIdx.tail) {
            //System.out.println(expect + ", " + i.toString)
            if (i._1 == expect) {
                return Array(numsIdx.head._2, i._2)
            }
        }
        return func(numsIdx.tail, target)
    }
}
```
```Scala
object Solution {
     def twoSum(nums: Array[Int], target: Int): Array[Int] = {
        val map = scala.collection.mutable.Map[Int, Int] ()
        var count = 0;
        for(num <- nums){
            val comp = target-num
            map.get(comp) match {
                case None => {map.put(num,count)}
                case Some(index) => return Array(index,count) 
            }
            count = count+1
        }
        Array(0,0)
    }
}
```