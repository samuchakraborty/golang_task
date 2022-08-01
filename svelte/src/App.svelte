<script lang="ts">
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
        <p> fetching </p>
    <!-- {:else if 5 > x}
        <p>{x} is less than 5</p> -->
    {:else}
        <p>{result.msg} </p>
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
