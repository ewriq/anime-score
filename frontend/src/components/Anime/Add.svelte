<script>
  import axios from "axios";

  let thumbnail = "";
  let name = "";
  let description = "";
  let err;

  async function add() {
    const form = {
      thumbnail: thumbnail,
      name: name,
      description: description,
    };

    try {
      const response = await axios.post(
        "http://localhost:3000/api/anime/add",
        form
      );
      console.error(response.data);
      if (response.data.status === "OK") {
        err = "İşlem başarılı";
        setTimeout(function () {
          window.location.href = "/admin";
        }, 3000);
      } else {
        err = "Anime zaten bulunuyor yada bir sorun var";
      }
    } catch (error) {
      console.error("Hata oluştu:", error);
    }
  }
</script>

<main>
  <div class="flex justify-start">
    <div class="w-full sm:w-96 p-4 ml-4">
      <form class="space-y-6" action="#">
        <h5 class="text-xl font-medium text-gray-900 dark:text-white">
          Anime ekle
        </h5>
        {#if err}
          <div class="bg-green-100 text-green-700 p-4" role="alert">
            <p class="font-bold">{err}</p>
          </div>
        {/if}
        <div>
          <label
            for="name"
            class="block mb-2 text-sm font-medium text-gray-900 dark:text-white"
            >Anime İsmi</label
          >
          <input
            type="name"
            bind:value={name}
            id="name"
            class="bg-gray-50 border border-gray-300 text-gray-900 text-sm rounded-lg focus:ring-blue-500 focus:border-blue-500 block w-full p-2.5 dark:bg-gray-600 dark:border-gray-500 dark:placeholder-gray-400 dark:text-white"
            required
          />
        </div>
        <div>
          <label
            for="thumbnail"
            class="block mb-2 text-sm font-medium text-gray-900 dark:text-white"
            >Küçük resim linki</label
          >
          <input
            type="thumbnail"
            bind:value={thumbnail}
            id="thumbnail"
            class="bg-gray-50 border border-gray-300 text-gray-900 text-sm rounded-lg focus:ring-blue-500 focus:border-blue-500 block w-full p-2.5 dark:bg-gray-600 dark:border-gray-500 dark:placeholder-gray-400 dark:text-white"
            required
          />
        </div>

        <div>
          <label
            for="description"
            class="block mb-2 text-sm font-medium text-gray-900 dark:text-white"
            >Açıklama kısmı</label
          >
          <input
            type="description"
            bind:value={description}
            id="description"
            class="bg-gray-50 border border-gray-300 text-gray-900 text-sm rounded-lg focus:ring-blue-500 focus:border-blue-500 block w-full p-2.5 dark:bg-gray-600 dark:border-gray-500 dark:placeholder-gray-400 dark:text-white"
            required
          />
        </div>

        <button
          on:click|preventDefault={add}
          type="submit"
          class="w-full text-white bg-blue-700 hover:bg-blue-800 focus:ring-4 focus:outline-none focus:ring-blue-300 font-medium rounded-lg text-sm px-5 py-2.5 text-center dark:bg-blue-600 dark:hover:bg-blue-700 dark:focus:ring-blue-800"
          >Ekle</button
        >
      </form>
    </div>
  </div>
</main>
