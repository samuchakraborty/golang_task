<script lang="ts">
    import { is_empty } from "svelte/internal";
    import BubbleSort from "./lib/BubbleSort.svelte";

    // import logo from "./assets/svelte.png";
    // import Counter from "./lib/Counter.svelte";
    let number;
    let result;

    let s: Array<number>;
    let bubbleSortTime: any;
    let insertSortTime: any;
    let mergeSortTime: any;

    async function calculateSorting() {
        console.log(number);
        const res = await fetch("http://localhost:8080/data", {
            method: "POST",
            body: JSON.stringify({
                number,
            }),
            headers: {
                "Content-Type": "application/json", //In my case i only used this header and it worked for me
            },
        });

        const json = await res.json();
        result = json;
        console.log(result);

        // let startTime = performance.now();
        var start = new Date().getTime();

        s = bubbleSort(result.usSort);
        // let endTime = performance.now()
        var end = new Date().getTime();

        bubbleSortTime = end - start;

        let startTimeInsertion = new Date().getTime();
        insertionSort(result.usSort);
        let endTimeInsertion = new Date().getTime();
        insertSortTime = endTimeInsertion - startTimeInsertion;

        let startTimeMerge = new Date().getTime();
        mergeSort(result.usSort);
        let endTimeMerge = new Date().getTime();
        mergeSortTime = endTimeMerge - startTimeMerge;
    }

    function bubbleSort(arr: Array<number>) {
        let array = arr;
        for (let i = 0; i < array.length; i++) {
            for (let j = 0; j < array.length - 1; j++) {
                if (array[j] > array[j + 1]) {
                    let num = array[j];
                    array[j] = array[j + 1];
                    array[j + 1] = num;

                    //fmt.Printf("%v %v \n", array[j], array[j+1])
                }
            }
        }
        return array;
    }
    function insertionSort(arr: Array<Number>) {
        let n = arr.length;
        for (let i = 1; i < n; i++) {
            let current = arr[i];
            let j = i - 1;
            while (j > -1 && current < arr[j]) {
                arr[j + 1] = arr[j];
                j--;
            }
            arr[j + 1] = current;
        }
        return arr;
    }

    function mergeSort(items: number[]): number[] {
        var halfLength = Math.ceil(items.length / 2);
        var low = items.slice(0, halfLength);
        var high = items.slice(halfLength);
        if (halfLength > 1) {
            low = mergeSort(low);
            high = mergeSort(high);
        }
        return combine(low, high);
    }

    function combine(low: number[], high: number[]): number[] {
        var indexLow = 0;
        var indexHigh = 0;
        var lengthLow = low.length;
        var lengthHigh = high.length;
        var combined = [];
        while (indexLow < lengthLow || indexHigh < lengthHigh) {
            var lowItem = low[indexLow];
            var highItem = high[indexHigh];
            if (lowItem !== undefined) {
                if (highItem === undefined) {
                    combined.push(lowItem);
                    indexLow++;
                } else {
                    if (lowItem <= highItem) {
                        combined.push(lowItem);
                        indexLow++;
                    } else {
                        combined.push(highItem);
                        indexHigh++;
                    }
                }
            } else {
                if (highItem !== undefined) {
                    combined.push(highItem);
                    indexHigh++;
                }
            }
        }
        return combined;
    }
</script>

<main class="text-center">
    <!-- <head>Samy</head> -->
    <div class="px-[30px] py-[30px] bg-indigo-500 my-20 mx-20 w-3/4">
        <label for="number">Enter your number :</label>
        <input type="text" id="number" name="number" bind:value={number} />
        <br />
        <br />
        <button
            class="bg-slate-400 p-4 rounded-full"
            on:click={calculateSorting}>Run</button
        >

        <!-- </form> -->
        <br />
    </div>
    <br />

    {#if result == null}
        <p>fetching</p>
        <!-- {:else if result }
        <p> is less than 5</p> -->
    {:else}
        <div class="overflow-x-auto relative mx-20">
            <p>GO executionTime</p>
            <table
                class="w-screen text-sm text-left text-gray-500 dark:text-gray-400"
            >
                <thead
                    class="text-xs text-gray-700 uppercase bg-gray-50 dark:bg-gray-700 dark:text-gray-400"
                >
                    <tr>
                        <th scope="col" class="py-3 px-6">
                            Bubble Sort ExecutionTime
                        </th>
                        <th scope="col" class="py-3 px-6">
                            Insertion Sort ExecutionTime
                        </th>
                        <th scope="col" class="py-3 px-6">
                            Merge Sort ExecutionTime
                        </th>
                    </tr>
                </thead>
                <tbody>
                    <tr class="bg-white dark:bg-gray-800">
                        <th
                            scope="row"
                            class="py-4 px-6 font-medium text-gray-900 whitespace-nowrap dark:text-white"
                        >
                            {result.bubble_sort_executionTime} miliseconds
                        </th>
                        <td class="py-4 px-6">
                            {result.insertion_sort_executionTime} miliseconds
                        </td>
                        <td class="py-4 px-6">
                            {result.merge_sort_executionTime} miliseconds
                        </td>
                    </tr>
                </tbody>
            </table>
        </div>
        <br />
        <br />
        <div class="overflow-x-auto relative mx-20">
            <p>TypeScript executionTime</p>

            <table
                class="w-screen text-sm text-left text-gray-500 dark:text-gray-400"
            >
                <thead
                    class="text-xs text-gray-700 uppercase bg-gray-50 dark:bg-gray-700 dark:text-gray-400"
                >
                    <tr>
                        <th scope="col" class="py-3 px-6">
                            Bubble Sort ExecutionTime
                        </th>
                        <th scope="col" class="py-3 px-6">
                            Insertion Sort ExecutionTime
                        </th>
                        <th scope="col" class="py-3 px-6">
                            Merge Sort ExecutionTime
                        </th>
                    </tr>
                </thead>
                <tbody>
                    <tr class="bg-white dark:bg-gray-800">
                        <th
                            scope="row"
                            class="py-4 px-6 font-medium text-gray-900 whitespace-nowrap dark:text-white"
                        >
                            {bubbleSortTime} miliseconds
                        </th>
                        <td class="py-4 px-6">
                            {insertSortTime} miliseconds
                        </td>
                        <td class="py-4 px-6">
                            {mergeSortTime} miliseconds
                        </td>
                    </tr>
                </tbody>
            </table>
        </div>
    {/if}
</main>
