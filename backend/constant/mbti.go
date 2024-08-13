package constant

type MBTI int

const (
	INTJ_A MBTI = iota // 0
	INTJ_T
	INTP_A
	INTP_T
	ENTJ_A
	ENTJ_T
	ENTP_A
	ENTP_T
	INFJ_A
	INFJ_T
	INFP_A
	INFP_T
	ENFJ_A
	ENFJ_T
	ENFP_A
	ENFP_T
	ISTJ_A
	ISTJ_T
	ISFJ_A
	ISFJ_T
	ESTJ_A
	ESTJ_T
	ESFJ_A
	ESFJ_T
	ISTP_A
	ISTP_T
	ISFP_A
	ISFP_T
	ESTP_A
	ESTP_T
	ESFP_A
	ESFP_T
)

func GetCompatibility(mbti MBTI) (goodCompatibility, badCompatibility []MBTI) {

	switch mbti {
	case INTJ_A, INTJ_T:
		goodCompatibility := []MBTI{ESFP_A, ESFJ_T, ISFJ_A, ISFJ_T, ENTJ_A, ENTJ_T}
		badCompatibility := []MBTI{ENTP_A, ENTP_T, ISFP_A, ISFP_T, ESFJ_A, ESFJ_T}
		return goodCompatibility, badCompatibility
	case INTP_A, INTP_T:
		goodCompatibility := []MBTI{ESFJ_A, ESFJ_T, ISFP_A, ISFP_T, ENTP_A, ENTP_T}
		badCompatibility := []MBTI{ENTJ_A, ENTJ_T, ISFJ_A, ISFJ_T, ESFP_A, ESFP_T}
		return goodCompatibility, badCompatibility
	case ENTJ_A, ENTJ_T:
		goodCompatibility := []MBTI{ISFJ_A, ISFJ_T, ESFP_A, ESFP_T, INTJ_A, INTJ_T}
		badCompatibility := []MBTI{INTP_A, INTP_T, ESFJ_A, ESFJ_T, ISFP_A, ISFP_T}
		return goodCompatibility, badCompatibility
	case ENTP_A, ENTP_T:
		goodCompatibility := []MBTI{ISFP_A, ISFP_T, ESFJ_A, ESFJ_T, INTP_A, INTP_T}
		badCompatibility := []MBTI{INTJ_A, INTJ_T, ESFP_A, ESFP_T, ISFJ_A, ISFJ_T}
		return goodCompatibility, badCompatibility
	case INFJ_A, INFJ_T:
		goodCompatibility := []MBTI{ESTP_A, ESTP_T, ISTJ_A, ISTJ_T, ENFJ_A, ENFJ_T}
		badCompatibility := []MBTI{ENFP_A, ENFP_T, ISTP_A, ISTP_T, ESTJ_A, ESTJ_T}
		return goodCompatibility, badCompatibility
	case INFP_A, INFP_T:
		goodCompatibility := []MBTI{ESTJ_A, ESTJ_T, ISTP_A, ISTP_T, ENFP_A, ENFP_T}
		badCompatibility := []MBTI{ENFJ_A, ENFJ_T, ISTJ_A, ISTJ_T, ESTP_A, ESTP_T}
		return goodCompatibility, badCompatibility
	case ENFJ_A, ENFJ_T:
		goodCompatibility := []MBTI{ISTJ_A, ISTJ_T, ESTP_A, ESTP_T, INFJ_A, INFJ_T}
		badCompatibility := []MBTI{INFP_A, INFP_T, ESTJ_A, ESTJ_T, ISTP_A, ISTP_T}
		return goodCompatibility, badCompatibility
	case ENFP_A, ENFP_T:
		goodCompatibility := []MBTI{ISTP_A, ISTP_T, ESTJ_A, ESTJ_T, INFP_A, INFP_T}
		badCompatibility := []MBTI{INFJ_A, INFJ_T, ESTP_A, ESTP_T, ISTJ_A, ISTJ_T}
		return goodCompatibility, badCompatibility
	case ISTJ_A, ISTJ_T:
		goodCompatibility := []MBTI{ENFJ_A, ENFJ_T, INFJ_A, INFJ_T, ESTP_A, ESTP_T}
		badCompatibility := []MBTI{ESTJ_A, ESTJ_T, INFP_A, INFP_T, ENFP_A, ENFP_T}
		return goodCompatibility, badCompatibility
	case ISFJ_A, ISFJ_T:
		goodCompatibility := []MBTI{ENTJ_A, ENTJ_T, INTJ_A, INTJ_T, ESFP_A, ESFP_T}
		badCompatibility := []MBTI{ESFJ_A, ESFJ_T, INTP_A, INTP_T, ENTP_A, ENTP_T}
		return goodCompatibility, badCompatibility
	case ESTJ_A, ESTJ_T:
		goodCompatibility := []MBTI{INFP_A, INFP_T, ENFP_A, ENFP_T, ISTP_A, ISTP_T}
		badCompatibility := []MBTI{ISTJ_A, ISTJ_T, ENFJ_A, ENFJ_T, INFJ_A, INFJ_T}
		return goodCompatibility, badCompatibility
	case ESFJ_A, ESFJ_T:
		goodCompatibility := []MBTI{INTP_A, INTP_T, ENTP_A, ENTP_T, ISFP_A, ISFP_T}
		badCompatibility := []MBTI{ISFJ_A, ISFJ_T, ENTJ_A, ENTJ_T, INTJ_A, INTJ_T}
		return goodCompatibility, badCompatibility
	case ISTP_A, ISTP_T:
		goodCompatibility := []MBTI{ENFP_A, ENFP_T, INFP_A, INFP_T, ESTJ_A, ESTJ_T}
		badCompatibility := []MBTI{ESTP_A, ESTP_T, INFJ_A, INFJ_T, ENFJ_A, ENFJ_T}
		return goodCompatibility, badCompatibility
	case ISFP_A, ISFP_T:
		goodCompatibility := []MBTI{ENTP_A, ENTP_T, INTP_A, INTP_T, ESFJ_A, ESFJ_T}
		badCompatibility := []MBTI{ESFP_A, ESFP_T, INTJ_A, INTJ_T, ENTJ_A, ENTJ_T}
		return goodCompatibility, badCompatibility
	case ESTP_A, ESTP_T:
		goodCompatibility := []MBTI{INFJ_A, INFJ_T, ENFJ_A, ENFJ_T, ISTJ_A, ISTJ_T}
		badCompatibility := []MBTI{ISTP_A, ISTP_T, ENFP_A, ENFP_T, INFP_A, INFP_T}
		return goodCompatibility, badCompatibility
	case ESFP_A, ESFP_T:
		goodCompatibility := []MBTI{INTJ_A, INTJ_T, ENTJ_A, ENTJ_T, ISFJ_A, ISFJ_T}
		badCompatibility := []MBTI{ISFP_A, ISFP_T, ENTP_A, ENTP_T, INTP_A, INTP_T}
		return goodCompatibility, badCompatibility
	}
	return
}
