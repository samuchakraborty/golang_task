<script lang="ts">



   export function BubbleSort(arr: Array<number>) {
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
</script>
