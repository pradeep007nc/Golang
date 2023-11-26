import java.util.HashMap;
import java.util.Map;
import java.util.Random;

public class TimeTaken {

    public static int[] twoSum(int[] nums, int target) {
        int[] ans = new int[2];
        Map<Integer, Integer> dictMap = new HashMap<>();

        for (int index = 0; index < nums.length; index++) {
            int val = nums[index];
            int compliment = target - val;
            if (dictMap.containsKey(compliment)) {
                ans[0] = dictMap.get(compliment);
                ans[1] = index;
                return ans;
            }
            dictMap.put(val, index);
        }

        return null;
    }

    public static int[] generateRandom(int arrSize, int limit) {
        int[] arr = new int[arrSize];
        Random random = new Random();

        for (int i = 0; i < arrSize; i++) {
            arr[i] = random.nextInt(limit + 1);
        }

        return arr;
    }

    public static void twoSumTime(int times, int arrSize) {
        for (int i = 0; i < times; i++) {
            int[] randomArr = generateRandom(arrSize, i);
            int target = new Random().nextInt(i + 1);
            twoSum(randomArr, target);
        }
    }

    public static void main(String[] args) {
        int[] result = twoSum(new int[]{2, 7, 11, 15}, 9);
        System.out.println(java.util.Arrays.toString(result));

        long startTime = System.currentTimeMillis();
        twoSumTime(100000, 100000);
        long endTime = System.currentTimeMillis();

        long total = endTime - startTime;

        System.out.println(total);
    }
}
