<script>
import { onMount } from "svelte";
let error, sucess;
// onMount(async () => {
//     try{
//         const config = {
//             headers: { Authorization: `Bearer ${Cookies.get('token')}` }
//         };
//         //get event data
//         const response = await axios.post(`/ticket/buy`, {
//             ticktId: ,
//         }, config);
//         res = response.data;
//         sucess = res;

//     }catch (e){
//         console.error("Error:", e);
//         error = response.data.error;
//     }
// })
</script>

<main class="flex col cen">
    <div class="margin-t">
        <h1>Yay you have successfuly bought a ticket! 😀 Enjoy the event!</h1>
    </div>
    {#if error != "" && error != undefined}
        <p class="error">{error}</p>
    {/if}
    {#if sucess != "" && sucess != undefined}
        <p class="success">{sucess}</p>
    {/if}
</main>

<style>
    h1{
        color: rgb(12, 84, 12);
        /* background-color: rgba(0, 87, 0, 0.731); */
        border-radius: 0.4rem;
    }
    .margin-t{
        margin-top: 5rem;
    }
</style>