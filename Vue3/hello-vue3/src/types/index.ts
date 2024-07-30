// 定义一个接口 用于限制Person对象的具体属性
export interface PersonInter {
    id: string
    name: string
    age: number
    x?:number   // '?'代表可有可无
}

// 一个自定义类型
export type Persons = Array<PersonInter>

