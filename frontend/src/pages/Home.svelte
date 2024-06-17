<script>
  import List from "../components/Anime/List.svelte";
  import VList from "../components/Vote/List.svelte";
  import Loading from "../components/Loading.svelte";
  export let users;
  import { onMount } from "svelte";
  import Search from "../components/Anime/Search.svelte";

  let loading = true;

  onMount(() => {
    setTimeout(() => {
      loading = false;
    }, 150);
  });
</script>

<div>
  {#if loading}
    <div class="flex justify-center min-h-screen items-center">
      <div
        class="mb-1.5 w-full max-w-sm p-4 bg-white border border-gray-200 rounded-lg shadow sm:p-6 md:p-8 dark:bg-gray-800 dark:border-gray-700"
      >
        <Loading />
        <h3
          class="text-2xl font-semibold whitespace-nowrap dark:text-white text-center"
        >
          Loading
        </h3>
      </div>
    </div>
  {:else if users}
    <div class="bg-white rounded-lg shadow m-4 dark:bg-gray-800">
      <Search  />
      <VList />
    </div>
  {:else}
    <main class="bg-white rounded-lg shadow m-4 dark:bg-gray-800">
      <Search  />
    </main>
  {/if}
</div>
