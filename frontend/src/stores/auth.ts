import { ref } from 'vue';
import { defineStore } from 'pinia';
import { connect, StringCodec } from "nats";

export const useAuthStore = defineStore({
    id: 'auth',
    state: () => ({}),
    actions: {
        async login(email: string, password: string) {
            const nc = await connect({
                servers: 'localhost',
                user: 'usr',
                pass: 'pass'
            })

            const res = await nc.request('login', StringCodec().encode(`{ email: ${email}, password: ${password} }`));
            console.log(res);

            return Promise.resolve();
        }
    },
});
    
    
    
//     'auth', async () => {
//     const nc = await connect({ servers: 'http://localhost:4222' });
//     const token = ref('');

//     function login(email: string, password: string) {
//         return new Promise((resolve) => {
//             const sc = StringCodec();
    
//             nc.request('login', sc.encode(`{ email: ${email}, password: ${password} }`))
//                 .then(data => {
//                     console.log(data);
//                     resolve([]);
//                 }).catch(err => {
//                     console.log(err);
//                     resolve([]);
//                 })
//         })

//     }

//     return { token, login };
// })