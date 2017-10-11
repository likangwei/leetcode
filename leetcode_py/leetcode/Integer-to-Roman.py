__author__ = 'likangwei'

class Solution:
    # @param {integer} num
    # @return {string}
    def intToRoman(self, num):
        m = 1
        result = ''
        while num/m > 0:
            cur_n = num%(m*10)/m*m
            # print cur_n
            result = self.get_one_num_string(cur_n) + result
            m = m * 10

        return result

    def get_one_num_string(self, num):
        if num is 0:
            return ''
        x = {'I': 1, 'V': 5, 'X': 10, 'L': 50, 'C': 100, 'D': 500, 'M': 1000}
        x_l = [('I', 1), ('V', 5), ('X', 10), ('L', 50), ('C', 100), ('D', 500), ('M', 1000)]

        cur = 0
        cur_c = None
        is_4_or_9 = str(num).startswith('4') or str(num).startswith('9')
        cur_n = 0
        if not is_4_or_9:
            for c,n in x_l:
                if num < n:
                    break
                else:
                    cur_n = n
                    cur_c = c
            num = num - cur_n
            if num:
                return cur_c + self.get_one_num_string(num)
            return cur_c
        else:
            for c,n in x_l:
                if n < num:
                    cur_n = n
                    cur_c = c
                else:
                    cur_n = n
                    cur_c = c
                    break

            num = cur_n - num
            if num:
                return self.get_one_num_string(num) + cur_c
            return cur_c



print Solution().intToRoman(20)

