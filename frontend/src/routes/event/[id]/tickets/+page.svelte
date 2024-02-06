<script>
    import { onMount } from "svelte";
    import axios from "axios";
    import "iconify-icon";
    export let data;
    let res;
    onMount(async () => {
    try{
        //get event data
        const response = await axios.get(`http://localhost:8000/event/${data.data}/tickets`);
        res = response.data;
        // title = res.eventName;
        // description = res.eventDescription;
        // category = res.eventCategory;
        // date = res.eventDate;
        // location = res.eventLocation;
        // cover = res.eventImage;
        // banner = res.eventBanner;
        // capacity = res.eventCapacity;
        // seatsRequired = Boolean(res.SeatsRequired);
        console.log(res);
    }catch (e){
        console.error("Error:", e);
    }
})
</script>
<main class="flex col cen">
    {#if res != "" && res != undefined}
    <div class="margin-t">
        {#each res as ticket}
            <div class="ticket">
                {ticket.EventName}
                {ticket.TicketDate}
                {ticket.TicketLocation}
                <a href={"/ticket/buy/"+ ticket.Id}><iconify-icon icon="mdi:cart-variant"  style="color: black"></iconify-icon></a>
            </div>
        {/each}
    </div>
    {/if}
</main>
<!-- <p>Show all avalible tickets for this event and a shopping cart button. With on click redirect to stripe and an sucess get user id from jwd and add to db.</p> -->

<style>
    .margin-t{
        margin-top: 5rem;
    }
    .ticket{
        background-color: var(--bg);
        margin: 1rem;
        padding: 1rem;
        border-radius: 0.4rem;
    }
</style>