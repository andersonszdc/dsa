const { MaxPriorityQueue } = require('@datastructures-js/priority-queue')

function findRelativeRanks(score) {
    const MaxHeap = new MaxPriorityQueue()


    for (const s of score) {
        MaxHeap.push(s)
    }

    const labels = ['Gold Medal', 'Silver Medal', 'Bronze Medal', '4', '5']
    const result = []

    for (let i = 0; i < score.length; i++) {
        const val = MaxHeap.pop()
        const idx = score.findIndex((v) => v === val)
        result[idx] = labels[i]
    }

    return result
}

findRelativeRanks([1, 4, 2, 5, 3])