<script setup lang="ts">
import { onMounted, ref } from 'vue';

interface User {
    id: number;
    firstname: string;
    lastname: string;
    email: string;
}

const users = ref<User[]>([])

async function getUsers() {
    const res = await fetch("/api/users")
    const data = await res.json()
    users.value = data as User[]
}

async function deleteUser(userId: number) {
    await fetch("/api/users/" + userId, {
        method: "DELETE"
    })

    await getUsers();
}

onMounted(() => {
    getUsers()
})
</script>

<template>
    <div class="container mx-auto">
        <div class="mt-5">
            <h3 class="text-xl">Users List</h3>
        </div>

        <div class="mt-5 border rounded-box border-base-content/5">
            <table class="table">
                <thead>
                    <tr>
                        <th>First Name</th>
                        <th>Last Name</th>
                        <th>Email</th>
                        <th align="right"></th>
                    </tr>
                </thead>
                <tbody>
                    <tr v-for="value in users">
                        <td>{{ value.firstname }}</td>
                        <td>{{ value.lastname }}</td>
                        <td>{{ value.email }}</td>
                        <td class="flex justify-end">
                            <div class="flex gap-3" >
                                <button class="btn btn-xs btn-info">Edit</button>
                                <button class="btn btn-xs btn-error" @click="deleteUser(value.id)">Delete</button>
                            </div>
                        </td>
                    </tr>
                </tbody>
            </table>
        </div>
    </div>
</template>
