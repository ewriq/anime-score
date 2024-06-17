<script>
  export let users;
  console.log(users);
  import { onMount } from "svelte";
  import Add from "../components/Anime/Add.svelte";
  import Del from "../components/Anime/Del.svelte";
  import List from "../components/Anime/List.svelte";

  let current = "Listele";

  const showPage = (page) => {
    current = page;
  };

  onMount(() => {
    showPage("Listele");
  });
</script>

<main class="bg-white rounded-lg shadow m-4 dark:bg-gray-800">
  <div
    class="w-full mx-auto max-w-screen-xl p-4 flex items-center justify-between space-x-4"
  >
     <span
      on:click={() => showPage("Listele")}
      class="text-sm text-gray-500 sm:text-center dark:text-gray-400 cursor-pointer"
      >Listele
    </span>
    <span
      on:click={() => showPage("Ekle")}
      class="text-sm text-gray-500 sm:text-center dark:text-gray-400 cursor-pointer"
      >Ekle</span
    >
    <span
      on:click={() => showPage("Sil")}
      class="text-sm text-gray-500 sm:text-center dark:text-gray-400 cursor-pointer"
      >Sil
    </span>
  </div>

  {#if current === "Listele"}
    <List />
  {:else if current === "Ekle"}
    <Add />
  {:else if current === "Sil"}
    <Del />
  {/if}
</main>
