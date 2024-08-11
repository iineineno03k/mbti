// TypeScriptでの型定義
export interface User {
    id: number;
    lastName: string;
    firstName: string;
    nickname: string;
    mbti: number;
}

const mbtiTypes = [
    "INTJ-A", "INTJ-T", "INTP-A", "INTP-T", "ENTJ-A", "ENTJ-T", "ENTP-A", "ENTP-T",
    "INFJ-A", "INFJ-T", "INFP-A", "INFP-T", "ENFJ-A", "ENFJ-T", "ENFP-A", "ENFP-T",
    "ISTJ-A", "ISTJ-T", "ISFJ-A", "ISFJ-T", "ESTJ-A", "ESTJ-T", "ESFJ-A", "ESFJ-T",
    "ISTP-A", "ISTP-T", "ISFP-A", "ISFP-T", "ESTP-A", "ESTP-T", "ESFP-A", "ESFP-T"
];

export const getMbtiType = (index: number): string => {
    return mbtiTypes[index] || "Unknown";
};
