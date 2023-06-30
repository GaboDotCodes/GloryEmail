<script setup>
import { computed, ref, watch } from 'vue';

const props = defineProps({
  email: Object,
});

const date = ref(props.email?._source?.date);
const subject = ref(props.email?._source?.subject);
const fromName = ref(props.email?._source?.xfrom);
const fromEmail = ref(props.email?._source?.from);
const toNames = ref(props.email?._source?.xto);
const toEmails = ref(props.email?._source?.to);
const body = ref(props.email?._source?.body);

watch(() => props.email, (email) => {
  date.value = email._source?.date;
  subject.value = email._source?.subject;
  fromName.value = email._source?.xfrom;
  fromEmail.value = email._source?.from;
  toNames.value = email._source?.xto;
  toEmails.value = email._source?.to;
  body.value = email._source?.body;
})

const removeTags = (stringWithTags) => stringWithTags?.replace(/<[^>]+>/g, "")

const getFromName = computed(() => {
  return removeTags(fromName.value)
})

const getFromEmail = computed(() => {
  return removeTags(fromEmail.value)
})

const getToNames = computed(() => {
  return toNames.value?.map((name) => removeTags(name));
})

const getToEmails = computed(() => {
  return toEmails.value?.map((email) => removeTags(email));
})

defineEmits(['backToList'])
</script>

<template>
  <div class="flex flex-col gap-y-3">
    <span class="material-symbols-outlined cursor-pointer" @click="$emit('backToList')">arrow_back</span>
    <span class="font-bold text-sm">{{ date }}</span>
    <h1 class="font-bold text-lg">{{ subject }}</h1>
    <span class="pl-4"><span class="font-bold">From: </span>{{ getFromName }} - {{ getFromEmail }}</span>
    <span class="pl-4"><span class="font-bold">To: </span>{{ getToNames?.[0] }} - {{ getToEmails?.[0] }}</span>
    <span class="text-lg font-bold mt-4">Message</span>
    <p class="whitespace-pre-wrap pl-4">{{ body }}</p>
  </div>
</template>