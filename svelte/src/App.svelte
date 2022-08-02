<script lang="ts">
import { is_empty } from "svelte/internal";


    // import logo from "./assets/svelte.png";
    // import Counter from "./lib/Counter.svelte";
    let number;
    let result;
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
    }

    function BubbleSort(arr: Array<number>)  {
	let array = arr
	for (let i = 0; i < array.length; i++) {
		for(let j = 0; j < array.length-1; j++ ){
			if (array[j] > array[j+1] ){
				let num = array[j]
				array[j] = array[j+1]
				array[j+1] = num

				//fmt.Printf("%v %v \n", array[j], array[j+1])

			}

		}
	}
	return array
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
                        <th scope="col" class="py-3 px-6"> Sorted Data </th>
                    </tr>
                </thead>
                <tbody>
                    <tr class="bg-white dark:bg-gray-800">
                        <th
                            scope="row"
                            class="py-4 px-6 font-medium text-gray-900 whitespace-nowrap dark:text-white"
                        >
                            {result.bubble_sort_data.executionTime}
                        </th>
                        <td class="py-4 px-6">
                            {result.insertion_sort_data.executionTime}
                        </td>
                        <td class="py-4 px-6">
                            {result.merge_sort_executionTime}
                        </td>
                        <td class="py-4 px-6 h-auto ">
                            {result.merge_sort_executionData}
                        </td>
                    </tr>
                </tbody>
            </table>
        </div>

        <!-- <div class="p-10">

    <table class="table-auto">

      
        <tr>
          <th>Bubble Sort ExecutionTime</th>
          <th>Insertion Sort ExecutionTime</th>
          <th>Merge Sort ExecutionTime</th>
          <th>Sorted Data</th>

        </tr>
        <tr>
          <td>{result.bubble_sort_data.executionTime}</td>
          <td>{result.insertion_sort_data.executionTime}</td>
          <td>{result.merge_sort_executionTime}</td>
          <td>{result.merge_sort_executionData}</td>
        </tr>

      </table> 
        </div> -->
    {/if}
</main>

<style>
    a {
        @apply text-blue-700;
    }
    a:hover {
        @apply underline;
    }
</style>
