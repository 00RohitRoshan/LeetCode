class Solution:
    def myPow(self, x: float, n: int) -> float:
        if n == 0 :
            return 1
        if n < 0 :
            x = 1.0/x
            n *= -1
        if n % 2 :
            r = self.myPow(x,n//2)
            return r*r*x
        else:
            r = self.myPow(x,n/2)
            return r*r
        