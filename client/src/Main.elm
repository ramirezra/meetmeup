module Main exposing (main)

import Api.Object.User
import Api.Query as Query
import Api.Scalar exposing (Id(..))
import Api.ScalarCodecs exposing (Id)
import Browser
import Graphql.Http
import Graphql.Operation exposing (RootQuery)
import Graphql.SelectionSet as SelectionSet exposing (SelectionSet)
import Html exposing (..)
import PrintAny
import RemoteData exposing (RemoteData)



-- MODEL


type alias Model =
    RemoteData (Graphql.Http.Error Response) Response


init : Flags -> ( Model, Cmd Msg )
init _ =
    ( RemoteData.Loading, makeRequest (Id "3") )


type alias Flags =
    ()


type alias Response =
    { username : String
    , id : Id
    }


query : Id -> SelectionSet Response RootQuery
query id =
    Query.user { id = id } <|
        SelectionSet.map2 Response
            Api.Object.User.username
            Api.Object.User.id


makeRequest : Id -> Cmd Msg
makeRequest id =
    id
        |> query
        |> Graphql.Http.queryRequest "http://100.119.96.75:8080/query"
        |> Graphql.Http.send (RemoteData.fromResult >> GotResponse)


main : Program Flags Model Msg
main =
    Browser.element
        { init = init
        , view = view
        , update = update
        , subscriptions = subscriptions
        }



--UPDATE


type Msg
    = GotResponse Model


update : Msg -> Model -> ( Model, Cmd Msg )
update msg model =
    case msg of
        GotResponse response ->
            ( response, Cmd.none )



-- VIEW


view : Model -> Html Msg
view model =
    div []
        [ Html.input [] []
        , div
            []
            [ PrintAny.view model ]
        ]



-- SUBSCRIPTIONS


subscriptions : Model -> Sub Msg
subscriptions model =
    Sub.none
