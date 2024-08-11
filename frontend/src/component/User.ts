// TypeScriptでの型定義
export interface User {
    id: number;
    lastName: string;
    firstName: string;
    nickname: string;
    mbti: number;
}

const mbtiTypes = [
    "建築家(INTJ-A)", "建築家(INTJ-T)", "論理学者(INTP-A)", "論理学者(INTP-T)", 
    "指揮官(ENTJ-A)", "指揮官(ENTJ-T)", "討論者(ENTP-A)", "討論者(ENTP-T)",
    "提唱者(INFJ-A)", "提唱者(INFJ-T)", "仲介者(INFP-A)", "仲介者(INFP-T)", 
    "主人公(ENFJ-A)", "主人公(ENFJ-T)", "運動家(ENFP-A)", "運動家(ENFP-T)",
    "管理者(ISTJ-A)", "管理者(ISTJ-T)", "擁護者(ISFJ-A)", "擁護者(ISFJ-T)", 
    "幹部(ESTJ-A)", "幹部(ESTJ-T)", "領事(ESFJ-A)", "領事(ESFJ-T)",
    "巨匠(ISTP-A)", "巨匠(ISTP-T)", "冒険家(ISFP-A)", "冒険家(ISFP-T)", 
    "起業家(ESTP-A)", "起業家(ESTP-T)", "エンターテイナー(ESFP-A)", "エンターテイナー(ESFP-T)"
];

export const getMbtiType =(index: number): string => {
    return mbtiTypes[index] || "Unknown";
};
