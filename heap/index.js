class MinHeap {
    constructor(maxSize) {
        this.arr = new Array(maxSize).fill(null)
        this.maxSize = maxSize
        this.heapSize = 0
    }

    _parent(i) {
        return Math.floor((i - 1) / 2)
    }

    _lChild(i) {
        return 2 * i + 1
    }

    _rChild(i) {
        return 2 * i + 2
    }

    _minHeapify(i) {
        const l = this._lChild(i)
        const r = this._rChild(i)

        let smallest = i

        if (l < this.heapSize && this.arr[l] < this.arr[smallest]) {
            smallest = l
        }

        if (r < this.heapSize && this.arr[r] < this.arr[smallest]) {
            smallest = r
        }

        if (smallest !== i) {
            const temp = this.arr[i]
            this.arr[i] = this.arr[smallest]
            this.arr[smallest] = temp
            this._minHeapify(smallest)
        }
    }

    poll() {
        if (this.heapSize <= 0) {
            return null
        }

        const root = this.arr[0]
        this.arr[0] = this.arr[this.heapSize - 1]
        this.arr[this.heapSize - 1] = null
        this.heapSize -= 1

        this._minHeapify(0)

        return root
    }

    add(x) {
        if (this.heapSize === this.maxSize) {
            return
        }

        let i = this.heapSize
        this.arr[i] = x
        this.heapSize += 1

        while (i !== 0 && this.arr[this._parent(i)] > this.arr[i]) {
            const temp = this.arr[i]
            this.arr[i] = this.arr[this._parent(i)]
            this.arr[this._parent(i)] = temp
            i = this._parent(i)
        }
    }

    size() {
        return this.heapSize
    }

    getMin() {
        return this.arr[0]
    }
}

function FindKLargestElements(arr, k) {
    const minHeap = new MinHeap(k)

    for (const element of arr) {
        if (element > minHeap.getMin()) {
            if (minHeap.size() == k) {
                minHeap.poll()
            }

            minHeap.add(element)
        }
    }

    const result = []

    for (let i = 0; i < k; i++) {
        result[i] = minHeap.poll()
    }

    console.log(result)
}

FindKLargestElements([6, 2, 10, 5, 9, 50, 23, 21, 1, 55], 5)