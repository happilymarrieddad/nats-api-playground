<script lang="ts">
	import { onMounted } from 'vue';
	import { defineComponent } from 'vue';
	//import { useStore } from 'vuex';
	import { useRouter } from 'vue-router';

	export default defineComponent({
		setup() {
			//const store = useStore();
			const router = useRouter();

			// There's probably a lot better way to do this...
			onMounted(() => {
				//store.dispatch('logout');
				router.push({ name: 'login' });
			});

			return {};
		},
	});
</script>

<template>
	<div class="logout">
		<h1>Logging out...</h1>
	</div>
</template>