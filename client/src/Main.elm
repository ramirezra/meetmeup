module Main exposing (main)

import Api.Object exposing (User)
import Api.Query as Query
import Api.ScalarCodecs exposing (Id)
import Browser
import Graphql.Http
import Graphql.Operation exposing (RootQuery)
import Graphql.SelectionSet as SelectionSet exposing (SelectionSet)
import Html exposing (..)
import RemoteData exposing (RemoteData)


type alias Model =
    RemoteData (Graphql.Http.Error Response) Response


type alias Response =
    Id


makeRequest : Cmd Msg
makeRequest =
    Query.user (Query.UserRequiredArguments "1") Api.Object.User
        |> Graphql.Http.queryRequest "localhost:8080/graphql"
        |> Graphql.Http.send (RemoteData.fromResult >> GotResponse)


type Msg
    = GotResponse Model


init : Flags -> ( Model, Cmd Msg )
init _ =
    ( RemoteData.Loading, makeRequest )


type alias Flags =
    ()


main : Program Flags Model Msg
main =
    Browser.element
        { init = init
        , view = view
        , update = update
        , subscriptions = subscriptions
        }


view : Model -> Html Msg
view model =
    div [] []


update : Msg -> Model -> ( Model, Cmd Msg )
update msg model =
    case msg of
        GotResponse response ->
            ( response, Cmd.none )


subscriptions : Model -> Sub Msg
subscriptions model =
    Sub.none
