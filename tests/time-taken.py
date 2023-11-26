import random
from datetime import datetime

def two_sum(nums, target):
    ans = [0, 0]
    dict_map = {}

    for index, val in enumerate(nums):
        compliment = target - val
        if compliment in dict_map:
            ans[0] = dict_map[compliment]
            ans[1] = index
            return ans
        dict_map[val] = index

    return None

def generate_random(arr_size, limit):
    return [random.randint(0, limit + 1) for _ in range(arr_size)]

def two_sum_time(times, arr_size):
    for _ in range(times):
        random_arr = generate_random(arr_size, _)
        target = random.randint(0, _ + 1)
        two_sum(random_arr, target)

if __name__ == "__main__":
    print(two_sum([2, 7, 11, 15], 9))

    start_time = datetime.now()
    two_sum_time(10000, 10000)
    end_time = datetime.now()

    total = end_time - start_time

    print(type(total))
    print(total)
