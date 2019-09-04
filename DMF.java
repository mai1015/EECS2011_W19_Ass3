import java.util.Comparator;
import java.util.PriorityQueue;

public class DMF<E> {
    private PriorityQueue<E> maxheap;
    private PriorityQueue<E> minheap;
    private Comparator<? super E> comparator;

    public DMF(Comparator<? super E> comparatorMin, Comparator<? super E> comparatorMax) {
        this.maxheap = new PriorityQueue<>(comparatorMin);
        this.minheap = new PriorityQueue<>(comparatorMax);
        this.comparator = comparatorMin;
    }

    public void balance() {
        if (Math.abs(maxheap.size() - minheap.size()) > 1) {
            if (maxheap.size() > minheap.size()) {
                minheap.add(maxheap.poll());
            } else {
                maxheap.add(minheap.poll());
            }
        }
    }

    public void insert(E e) {
        if (isEmpty()) {
            minheap.add(e);
        } else if (comparator.compare(getMed(), e) < 0) {
            maxheap.add(e);
        } else {
            minheap.add(e);
        }
        balance();
    }

    public E getMed() {
        if (minheap.size() >= maxheap.size()) {
            return minheap.peek();
        }
        return maxheap.peek();
    }

    public E removeMed() {
        if (minheap.size() >= maxheap.size()) {
            return minheap.poll();
        }
        return maxheap.poll();
    }

    public int size() {
        return maxheap.size() + minheap.size();
    }

    public boolean isEmpty() {
        return maxheap.isEmpty() && minheap.isEmpty();
    }
}
