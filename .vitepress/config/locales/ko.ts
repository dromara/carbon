import {type DefaultTheme, defineConfig} from 'vitepress'

export const ko = defineConfig({
    lang: 'ko-KR',
    title: 'carbon',
    description: '가벼우면서도 의미론적이고 개발자 친화적인 golang 시간 처리 라이브러리',
    head: [
        ['meta', { name: 'keywords', content: 'golang, carbon, 시간 처리, 별자리, 계절, 그레고리력, 음력, 율리우스일, 단순 율리우스일, 페르시아력, 이란력' }],
    ],
    themeConfig: {
        nav: nav(),

        sidebar: {
            '/ko/': { base: '/ko/', items: sidebarGuide() },
        },

        editLink: {
            pattern: 'https://github.com/dromara/carbon/edit/docs/src/:path',
            text: 'GitHub에서 이 페이지 편집'
        },

        footer: {
            message: 'MIT 라이선스에 따라 배포되며, 허가 없이 어떤 형태로든 재배포를 금지합니다',
            copyright: `저작권 © 2020-${new Date().getFullYear()} carbon team`
        },

        docFooter: {
            prev: '이전 페이지',
            next: '다음 페이지'
        },

        outline: {
            level: [2, 6],
            label: '현재 페이지'
        },

        lastUpdated: {
            text: '마지막 업데이트',
            formatOptions: {
                dateStyle: 'short',
                timeStyle: 'medium'
            }
        },

        langMenuLabel: '언어',
        returnToTopLabel: '맨 위로',
        sidebarMenuLabel: '메뉴',
        darkModeSwitchLabel: '테마',
        lightModeSwitchTitle: '라이트 모드로 전환',
        darkModeSwitchTitle: '다크 모드로 전환'
    }
})

function nav(): DefaultTheme.NavItem[] {
    return [
        {
            text: '홈',
            link: '/ko'
        },
        {
            text: '사용 문서',
            link: '/ko/overview',
            activeMatch: '/ko/overview'
        },
        {
            text: '자주 묻는 질문',
            link: '/ko/faq',
            activeMatch: '/ko/faq'
        },
        {
            text: '업데이트 로그',
            link: '/ko/change-log',
            activeMatch: '/ko/change-log'
        },
        {
            text: '후원 지원',
            link: '/ko/sponsor',
            activeMatch: '/ko/sponsor'
        },
        {
            text: 'Dongle',
            link: 'https://dongle.go-pkg.com/',
        },
    ]
}

function sidebarGuide(): DefaultTheme.SidebarItem[] {
    return [
        {
            text: '시작 가이드',
            collapsed: false,
            items: [
                {text: '프로젝트 소개', link: 'overview',},
                {text: '빠른 시작', link: 'getting-started',},
            ]
        },
        {
            text: '기본 사용법',
            collapsed: false,
            items: [
                {text: '기본값 설정', link: 'usage/default'},
                {text: 'carbon/time.Time 상호 변환', link: 'usage/conversion'},
                {text: '어제/오늘/내일', link: 'usage/yesterday-today-tomorrow'},
                {text: 'Carbon 인스턴스 생성', link: 'usage/creator'},
                {text: '시간 파싱', link: 'usage/parser'},
                {text: '시간 여행', link: 'usage/traveler'},
                {text: '시간 설정', link: 'usage/setter'},
                {text: '시간 가져오기', link: 'usage/getter'},
                {text: '시간 출력', link: 'usage/output'},
                {text: '시간 경계', link: 'usage/boundary'},
                {text: '시간 차이', link: 'usage/difference'},
                {text: '시간 극값', link: 'usage/extremum'},
                {text: '시간 판단', link: 'usage/comparison'},
                {text: '오류 처리', link: 'usage/error'},
            ]
        },
        {
            text: '확장 사용법',
            collapsed: false,
            items: [
                {text: '별자리', link: 'usage/constellation'},
                {text: '계절', link: 'usage/season'},
                {text: '달력', link: 'usage/calendar'},
                {text: '국제화', link: 'usage/i18n'},
                {text: 'JSON', link: 'usage/json'},
                {text: '테스트', link: 'usage/test'},
            ]
        },
        {
            text: '부록',
            collapsed: false,
            items: [
                {text: '내장 상수', link: 'appendix/constants'},
                {text: '형식 기호', link: 'appendix/format-signs'},
                {base:'https://', text: '테스트 보고서', link: 'github.com/dromara/carbon/blob/master/test_report.ko.md', target: '_blank'},
            ]
        },
    ]
}

export const search: DefaultTheme.AlgoliaSearchOptions['locales'] = {
    ko: {
        placeholder: '문서 검색',
        translations: {
            button: {
                buttonText: '문서 검색',
                buttonAriaLabel: '문서 검색'
            },
            modal: {
                searchBox: {
                    resetButtonTitle: '검색 조건 초기화',
                    resetButtonAriaLabel: '검색 조건 초기화',
                    cancelButtonText: '취소',
                    cancelButtonAriaLabel: '취소'
                },
                startScreen: {
                    recentSearchesTitle: '검색 기록',
                    noRecentSearchesText: '검색 기록이 없습니다',
                    saveRecentSearchButtonTitle: '검색 기록에 저장',
                    removeRecentSearchButtonTitle: '검색 기록에서 제거',
                    favoriteSearchesTitle: '즐겨찾기',
                    removeFavoriteSearchButtonTitle: '즐겨찾기에서 제거'
                },
                errorScreen: {
                    titleText: '결과를 가져올 수 없습니다',
                    helpText: '네트워크 연결을 확인해보세요'
                },
                footer: {
                    selectText: '선택',
                    navigateText: '전환',
                    closeText: '닫기',
                    searchByText: '검색 제공자'
                },
                noResultsScreen: {
                    noResultsText: '관련 결과를 찾을 수 없습니다',
                    suggestedQueryText: '다음 검색어를 시도해보세요',
                    reportMissingResultsText: '이 검색어에 결과가 있어야 한다고 생각하시나요?',
                    reportMissingResultsLinkText: '클릭하여 피드백'
                }
            }
        }
    }
} 