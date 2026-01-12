<script setup lang="ts">
import { onMounted, ref } from 'vue';

interface Contact {
    id: number;
    name: string;
    email: string;
    phone_number: string;
    user_id: number
}

const contacts = ref<Contact[]>([])

async function getContacts() {
    const res = await fetch("/api/contacts?limit=500");
    const data = await res.json();

    contacts.value = data as Contact[]
}

async function deleteContact(contactId: number) {
    await fetch("/api/contacts/" + contactId, {
        method: "DELETE"
    })

    await getContacts();
}


onMounted(() => {
    getContacts()
})
</script>

<template>
    <div class="container mx-auto">
        <div class="mt-5">
            <h3 class="text-xl">Contacts List</h3>
        </div>

        <div class="mt-5 border rounded-box border-base-content/5">
            <table class="table">
                <thead>
                    <tr>
                        <th>Name</th>
                        <th>Email</th>
                        <th>Phone Number</th>
                        <th>User</th>
                        <th></th>
                    </tr>
                </thead>
                <tbody>
                    <tr v-for="value in contacts">
                        <td>{{ value.name }}</td>
                        <td>{{ value.email }}</td>
                        <td>{{ value.phone_number }}</td>
                        <td>{{ value.user_id }}</td>
                        <td class="flex gap-3 justify-end">
                            <button class="btn btn-info btn-xs">Edit</button>
                            <button class="btn btn-error btn-xs" @click="deleteContact(value.id)">Delete</button>
                        </td>
                    </tr>
                </tbody>
            </table>
        </div>
    </div>
</template>