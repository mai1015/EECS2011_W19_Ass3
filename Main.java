import java.util.Comparator;

public class Main {

    static class Ca implements Comparator<Integer> {

        @Override
        public int compare(Integer o1, Integer o2) {
            return o1 - o2;
        }
    }

    static class Cb implements Comparator<Integer> {

        @Override
        public int compare(Integer o1, Integer o2) {
            return o2 - o1;
        }
    }

    public static int[][] example = {
            {4, 9, 1}, {9, 3, 3}, {9, 9, 1, 2}, {17, -4, 13, -7, 13, 15, 5, 2}
    };

    public static void main(String[] args) {
        for (int i = 0; i < example.length; i++) {
            DMF<Integer> dmf = new DMF<>(new Ca(), new Cb());
            for (int j = 0; j < example[i].length; j++) {
                dmf.insert(example[i][j]);
            }
            System.out.println(dmf.getMed());
        }
    }
}
